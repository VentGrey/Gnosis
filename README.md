# 🧬 GNOSIS 👁

Gnosis is a very simple CLI tool written in Go designed to generate TypeScript definition files from PocketBase JSON outputs. Did I mention it's stdlib only? No dependencies!

PocketBase JSON outputs typically contain a list of items, each with a special field `@collectionName` that specifies the type of the item. Gnosis uses this field to determine the name of the TypeScript interface it generates.

That's pretty much it...I said it was simple. The codebase isn't large either.

> I've originally wrote this in Perl circa ~2022. I thought it needed a "remasterization" so I rewrote it in Go.

## 🚀 Installation

1. Clone the repository:

```sh
$ git clone https://github.com/your_username/gnosis.git
$ cd gnosis
```

2. Build the project using the provided Makefile:

```sh
$ make

# Or

$ make build
```

The Makefile has the following targets:

| Target | Description |
| ------ | ----------- |
| `all` or `build` | Builds the project |
| `install` | Installs the binary to your system |
| `clean` | Cleans the project directory |
| `uninstall` | Uninstalls the binary from your system |
| `test` | Runs the tests |
| `help` | Prints the help message |

3. Install the binary to your system:

```sh
$ make install
```

This will install the `gnosis` binary to the default directory (`/usr/local/bin`). Ensure this directory is in your system's `$PATH` to run the `gnosis` command from anywhere.

## 📖 Usage

You can specify the input and output files using the -i and -o options:

```sh
$ gnosis -i input.json -o output.d.ts
```


## 🤔 Example

If you want a straightforward example, use the provided `example.json` file:

```sh
$ gnosis -i example.json
```

Otherwise, here's a more detailed example.

Given a PocketBase JSON output from a collection named "bands", which looks like this:

``` json
{
    "page": 1,
    "perPage": 200,
    "totalItems": 3,
    "totalPages": 1,
    "items": [
        {
            "@collectionId": "xt2edmu0xn980wn",
            "@collectionName": "bands",
            "name": "Opeth",
            "genre": "Progressive Metal",
            "origin": "Sweden",
            "formed": 1990,
            ...
        },
        {
            "@collectionId": "yt3epo1yn781zp",
            "name": "Alcest",
            "genre": "Post-Black Metal",
            "origin": "France",
            "formed": 2000,
            ...
        },
        {
            "@collectionId": "zp4epa2zt891qo",
            "name": "Riverside",
            "genre": "Progressive Rock",
            "origin": "Poland",
            "formed": 2001,
            ...
        },
        ...
    ]
}
```

Running `gnosis` providing the input file will generate a TypeScript definition file with the name of the collection:

``` sh
# Assuming the previous JSON is in a file called pocketbase.json (Name doesn't matter)
$ gnosis -i pocketbase.json
```

`gnosis` will analyze the JSON, particularly the first item in the items array to determine the structure. It will then generate and print an interface named after the `@collectionName` field value:

```typescript
interface Bands {
    name: string;
    genre: string;
    origin: string;
    formed: number;
    ...
}
```

With this interface, you can now have type-safe structures in your TypeScript project that match your PocketBase data.

If you don't specify the `-o` option, `gnosis` will print the TypeScript definition directly to the console. You can use this to pipe the output to a file or automate the typescript definition generation process.

```sh
$ gnosis -i input.json > output.d.ts

# For example, if you have a folder containing a bunch of JSON files, you can generate the TypeScript definitions for all of them like this:

$ for file in *.json; do gnosis -i $file > $file.d.ts; done

# Another example using curl to automatically generate TypeScript definitions for a PocketBase collection taken from a URL (Assuming you have a PocketBase URL to fetch the JSON from):

$ curl -s https://pocketbase.com/api/collections/1234 | gnosis > collection.d.ts
```

## 📝 Contributions

Contributions are welcome. Please open an issue first to discuss what you'd like to change.

## 📜 License

This project is licensed under the Gnu General Public License v2.0 ONLY - see the [LICENSE](LICENSE) file for details.

## 📜 Disclaimer

The bands and artists mentioned in the examples provided in this repository are purely for illustrative purposes. The author has no intention of infringing any copyrights. The references to these bands and artists are merely indicative of the author's personal taste in music and are not intended for commercial use or misrepresentation. All rights to the names and trademarks of the bands and artists belong to their respective owners.
