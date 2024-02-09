const fs = require('fs');
const axios = require('axios');
const path = require('path');
const archiver = require('archiver');
const moment = require('moment');

async function fetchSeriesValues(inputFile) {

    const regex = /(#\d+) - (.+?) - (.+?) (\d+ ka)/;
    const inputData = fs.readFileSync(inputFile, 'utf-8');
    const _data = inputData.split('\n');
    const _head = _data.slice(0, 3);
    const _body = _data.slice(4);

    const title = _head[0].trim();
    const total = _head[2].match(/\d+/)[0];

    const data = {
	metadata: {
	    title: title,
	    total: total
	}
    };
  
    const characters = [];

    for (const line of _body) {
	// Line is empty, skip it
	if (!line.trim()) {
	    continue;
	}

	const elements = line.match(regex);
      
	if (elements) {
	    const character = {};
	    character.rank = elements[1];
	    character.name = elements[2];
	    character.series = elements[3];
	    character.value = elements[4];
	    characters.push(character);
	} else {
	    console.error(`Error parsing line: ${line}`);
	}
    }

    data.characters = characters;
    return data;
}

async function fetchNotesImages(inputFile, data) {

    const regex = /^(#\d+).{3}([^|]+)(?: \| (.+)).{2} (.+)/;
    const inputData = fs.readFileSync(inputFile, 'utf-8');
    const _data = inputData.split('\n');
    // const _head = _data.slice(0, 1);
    const _body = _data.slice(2);

    for (const line of _body) {
	// Line is empty, skip it
	if (!line.trim()) {
	    continue;
	}

	const elements = line.match(regex);
      
	if (elements) {
	    var result = data.characters.find(search => search.name === elements[2]);

	    result.note = elements[3] || '';
	    result.image = elements[4];
	} else {
	    console.error(`Error parsing line: ${line}`);
	}
    }
}


async function saveJsonToFile(data, filename) {
    const jsonFilename = filename;
    fs.writeFileSync(jsonFilename, JSON.stringify(data, null, 2));
    console.log(`JSON data saved to ${jsonFilename}`);
}

async function downloadImages(characters, imagesDirectory) {

    if (!characters || !characters.length) {
	console.error('No data to download images.');
	return;
    }

    try {
	// Check if imagesDirectory exists, if not, create it
	if (!fs.existsSync(imagesDirectory)) {
	    fs.mkdirSync(imagesDirectory, { recursive: true });
	}

	for (const [index, character] of characters.entries()) {
	    const imageUrl = character.image;
	    const rank = character.rank.match(/\d+/)[0];
	    const name = character.name.replace(/\s+/g, '_');
	    const extension = imageUrl.substring(imageUrl.lastIndexOf('.') + 1);
	    const localImagePath = path.join(imagesDirectory, `${index + 1}_${rank}_${name}.${extension}`);

	    try {
		const response = await axios.get(imageUrl, { responseType: 'arraybuffer' });
		fs.writeFileSync(localImagePath, Buffer.from(response.data));
		console.log(`Image for ${character.name} downloaded and saved to ${localImagePath}`);
	    } catch (error) {
		console.error(`Error downloading image for ${character.name}: ${error.message}`);
	    }
	}
    } catch (error) {
	console.error(`Error creating images directory: ${error.message}`);
    }
}

async function replaceRemotePathsWithLocal(data, imagesDirectory) {

    data.characters.forEach((character, index) => {
	const rank = character.rank.match(/\d+/)[0];
	const name = character.name.replace(/\s+/g, '_');
	const extension = character.image.substring(character.image.lastIndexOf('.') + 1);
	const localImagePath = `images/${index + 1}_${rank}_${name}.${extension}`;
	character.image = localImagePath;
    });

    await saveJsonToFile(data, 'data.json');
}

async function createZip(exportsDirectory, jsonFilename, imagesDirectory) {

    // Check if imagesDirectory exists, if not, create it
    if (!fs.existsSync(exportsDirectory)) {
	fs.mkdirSync(exportsDirectory, { recursive: true });
    }

    const data = require('./' + jsonFilename);
    const archive = archiver('zip', { zlib: { level: 9 } });
    const zipFilename = exportsDirectory + `export_${moment().format('YYYYMMDD_HHmmss')}.zip`;
    const zipStream = fs.createWriteStream(zipFilename);

    archive.pipe(zipStream);

    // Add the textJson as a file in the zip
    archive.append(JSON.stringify(data, null, 2), { name: jsonFilename });

    // Add images to the zip
    archive.directory(imagesDirectory, imagesDirectory);

    // Finalize the zip creation
    await archive.finalize();

    console.log(`Zip file ${zipFilename} created successfully.`);

    // Delete data.json and images/ after .zip creation
    await fs.rmSync(imagesDirectory, { recursive: true });
    await fs.unlinkSync(jsonFilename);
}

// Example usage, replace variables if needed
async function main() {

    try {
	const dataJson = await fetchSeriesValues('output1_series_values.txt');
	await fetchNotesImages('output2_notes_images.txt', dataJson);
	await saveJsonToFile(dataJson, 'data.json');
	await downloadImages(dataJson.characters, 'images/');
	await replaceRemotePathsWithLocal(dataJson, 'images/');
	await createZip('exports/', 'data.json', 'images/');
    } catch (error) {
	console.error('An error occurred:', error);
    }
}

main();
