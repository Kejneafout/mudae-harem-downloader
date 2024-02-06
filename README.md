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

- Run the `$mmi-r-k-s` command on your favorite server,
- Copy the entire output, from the harem title, to the end,
- Paste it in `output.txt`.

I left my perfect harem in the `output.txt` file as an example.
So you know what to paste.

### 2. Run the Script

Run the script using Node.js. Execute the following command in the terminal:

```bash
node index.js
```

This will download character images, create a `.json` file, and create a `.zip` archive containing the `.json` file and `images`.

`export_YYYYMMDD_HHmmss.zip`
|- `data.json`
|- `images/`
  |- ...
  |- ...

Images are named as follows:
`[index]_[rank]_[name].{png,gif}`

- `[index]` is the order in which you sorted your harem with $sm and $smp,
- `[rank]` is the character's claim rank,
- `[name]` is the character's name.

### 3. Access the Output

Once the script has finished running, you will find the following files in the `exports/` directory:

- `export_<timestamp>.zip`: archive containing `data.json` and character `images/`.

### 4. Limitations

- Does not take alias2 into account, only notes,
- Does not download images in formats other than .png or .gif, such as .webp,
- Does not take like ranks into account, only claim ranks (`r-` flag),
- Does not take keys into accounts (`k-` flag).

## License

This project is licensed under the GNU General Public License (GPL) - see the [LICENSE](LICENSE) file for details.
