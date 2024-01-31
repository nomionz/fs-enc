# FSENC
FSENC is a lightweight, command-line interface tool designed for encrypting and decrypting files using a passphrase. It's an ideal solution for those who need to secure individual files containing sensitive information without the overhead of encrypting their entire filesystem.

### Download from [Releases](https://github.com/nomionz/fsenc/releases/)

## Why FSENC?

- Selective Encryption: Perfect for encrypting specific files in your home directory or any other location, ensuring sensitive information is kept private.
- Ease of Use: A straightforward CLI tool, making file encryption and decryption as simple as a single command.
- Security: Uses robust encryption algorithms to secure your files, safeguarding them against unauthorized access.

## Usage

## Encrypting a File
To encrypt a file, use the enc command followed by the file path and a passphrase of your choice:
```bash
fsenc enc <filepath> <passphrase>
```

## Decrypting a File
To decrypt an encrypted file, use the dec command followed by the file path and the correct passphrase:
```bash
fsenc dec <filepath> <passphrase>
```
