# Mudae Harem Downloader

Mudae Harem Downloader is a tool for converting data from the Mudae game into a ZIP archive containing JSON data and character images.

That archive can then be used with [Mudae Harem Viewer](https://github.com/Kejneafout/mudae-harem-viewer) to view your harem as if you typed the `$mm` command !

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [How it works](#how-it-works)
- [License](#license)

## Installation

You have two options:
1. Download and use the pre-compiled binaries in [Releases](https://github.com/Kejneafout/mudae-harem-downloader/releases/tag/v1.0.0)
2. Clone the repository and compile the app yourself:

- You need [Go](https://golang.org/doc/install) installed on your system.
- Clone this repository to your local machine.
   ```bash
   git clone https://github.com/Kejneafout/mudae-harem-downloader.git
   ```
- Navigate to the project directory.
   ```bash
   cd mudae-harem-downloader
   ```
- Build the executable.
   ```bash
   go build .
   ```

## Usage

### 1. Prepare input files

- Run the `$mmsr-a+ky+` and `$mmsi-` commands on your favorite server,
- Copy and paste the output of `$mmsr-a+ky+` in `inputs/1_keys_series_values.txt`,
- Copy and paste the output of `$mmsi-` in `inputs/2_notes_images.txt`.

**Info:** I left my soulmates list in both input files, as an example.

**Warnings:**
- Copy the outputs manually, **DO NOT** `Right-click => Copy text` !
- Some text editors do not like the special dot separator '·' used for keys, make sure your editor doesn't corrupt the data because of this !

### 2. Run the executable

- On Linux:
```bash
./mudae-harem-downloader-linux > logs.txt
```

- On Windows:
```cmd
.\mudae-harem-downloader-windows.exe > logs.txt
```

### 3. Access the output archive

Once the script has finished running, you will find the following files in the `exports/` directory:

- `export_YYYYMMDD_HHmmss.zip`: archive containing `data.json` and character `images/`.

You can upload this to my other tool: [Mudae Harem Viewer](https://github.com/Kejneafout/mudae-harem-viewer)

## How it works

The script:
- Converts data from both input files into JSON format and save it to `data.json`,
- Download remote `.png` or `.gif` images and save them in the `images/` directory,
- Change remote image paths to local image paths in `images/`,
- Create a `.zip` archive in the `exports/` directory, containing the `data.json` file and `images/` directory,
- Once the `.zip` is done, delete `data.json` and `images/`.

Images are named as follows:
`[index]_[rank]_[name].{png,gif}`

- `[index]` is the order in which you sorted your harem with `$sm` and `$smp`,
- `[rank]` is the character's claim rank,
- `[name]` is the character's name.

## Limitations

- Does not take support aliases (`$a2`), only notes (`$n`), still works if you don't have notes,
- Does not download images in formats other than `.png` or `.gif`, such as `.webp`,
- Does not take like ranks into account, only claim ranks (`r-` flag),
- Does not support custom embed colors (`$ec`).

## License

This project is licensed under the GNU General Public License (GPL) - see the [COPYING](COPYING) file for details.
