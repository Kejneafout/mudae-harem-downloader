# Mudae Harem Downloader

Mudae Harem Downloader is a tool for converting data from the Mudae game into a ZIP archive containing JSON data and character images.

That archive can then be used with [Mudae Harem Viewer](https://github.com/Kejneafout/mudae-harem-viewer) to view your harem as if you typed the `$mm` command !

## Installation

1. Install Node.js in your system, preferably using [nvm](https://github.com/nvm-sh/nvm) for Linux or [nvm-windows](https://github.com/coreybutler/nvm-windows) for Windows

2. Clone the repository:

   ```bash
   git clone https://github.com/Kejneafout/mudae-harem-downloader.git
   ```

3. Navigate to the project directory:

   ```bash
   cd mudae-harem-downloader
   ```

4. Install dependencies:

   ```bash
   npm install
   ```

## Dependencies

The Mudae Harem Downloader relies on the following npm packages:

- [Axios](https://www.npmjs.com/package/axios): Promise-based HTTP client for the browser and Node.js.
- [Archiver](https://www.npmjs.com/package/archiver): Streaming interface for creating and extracting zip archives.
- [Moment](https://www.npmjs.com/package/moment): Parse, validate, manipulate, and display dates and times in JavaScript.

## Usage

### 1. Prepare Input Data

- Run the `$mmsr-a+k-` and `$mmsr-i-` commands on your favorite server,
- Copy and paste the output of `$mmsr-a+k-` in `output1_series_values.txt`.
- Copy and paste the output of `$mmsr-i-` in `output2_notes_images.txt`.

**Note:** Copy the outputs manually, do not Right-click => Copy text.

I left my perfect harem in both output files as an example.

### 2. Run the Script

Run the script using Node.js. Execute the following command in the terminal:

```bash
node index.js
```

The script will:
- Convert data from both output files into JSON format and save it to `data.json`,
- Download remote `.png` or `.gif` images and save them in the `images/` directory,
- Change remote image paths to local image paths in `images/`,
- Create a `.zip` archive in the `exports/` directory, containing the `data.json` file and `images/` directory,
- Once the `.zip` is done, delete `data.json` and `images/`.

Images are named as follows:
`[index]_[rank]_[name].{png,gif}`

- `[index]` is the order in which you sorted your harem with `$sm` and `$smp`,
- `[rank]` is the character's claim rank,
- `[name]` is the character's name.

### 3. Access the Output

Once the script has finished running, you will find the following files in the `exports/` directory:

- `export_YYYYMMDD_HHmmss.zip`: archive containing `data.json` and character `images/`.

You can upload this to my other tool: [Mudae Harem Viewer](https://github.com/Kejneafout/mudae-harem-viewer)

## Limitations

- Does not take `$a2` into account, only `$n`, it will still work if you don't have any,
- Does not download images in formats other than `.png` or `.gif`, such as `.webp`,
- Does not take like ranks into account, only claim ranks (`r-` flag),
- Does not take keys value into account (`k-` flag).

## License

This project is licensed under the GNU General Public License (GPL) - see the [LICENSE](LICENSE) file for details.
