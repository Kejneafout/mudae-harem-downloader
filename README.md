# Mudae Harem Downloader

Mudae Harem Downloader is a tool for converting data from the Mudae game into a .zip file containing JSON data and character images.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Kejneafout/mudae-harem-downloader.git
   ```

2. Navigate to the project directory:

   ```bash
   cd mudae-harem-downloader
   ```

3. Install dependencies:

   ```bash
   npm install
   ```

## Usage

### 1. Prepare Input Data

1. Run the $mmi-r-k-s command on your favorite server
2. Copy the entire output, from the harem title, to the end
3. Paste it in `output.txt`.

I left my perfect harem in the `output.txt` file as an example.
So you know what to paste.

Does not take alias2 into account, only notes.
If you have aliases2, change them to notes, and then run the $mmi-r-k-s command again.

### 2. Run the Script

Run the script using Node.js. Execute the following command in the terminal:

```bash
node index.js
```

This will download character images, create a JSON file, and create a .zip archive containing the JSON file and images.

### 3. Access the Output

Once the script has finished running, you will find the following files in the project directory:

- `data.json`: JSON file containing character data.
- `export_<timestamp>.zip`: .zip archive containing `data.json` and character images.

## License

This project is licensed under the GNU General Public License (GPL) - see the [LICENSE](LICENSE) file for details.