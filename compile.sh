#![/bin/sh]

abigen  --abi "./out/BotGuard.sol/BotGuard.abi.json" --type BotGuard --pkg binding --out binding/IBotGuard.go

abigen --abi "./contracts/seadrop/out/IERC721A.sol/IERC721A.abi.json" --type IERC721A --pkg binding --out binding/IERC721A.go

abigen --abi "./contracts/seadrop/out/ISeaDrop.sol/ISeaDrop.abi.json" --type ISeaDrop --pkg binding --out binding/ISeaDrop.go