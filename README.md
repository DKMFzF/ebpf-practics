# hello world on eBPF

guide: https://ebpf-go.dev/guides/getting-started/

## how to start?

**I use vm in yandex cloud instance with distro ubuntu 24.04 lts and version kernel 6.8.0-134-generic**

```bash
# critical deps
sudo apt install llvm clang linux-headers-generic linux-headers-$(uname -r) libbpf-dev 

sudo ln -sf /usr/include/asm-generic /usr/include/asm
# or if not started go generate use:
#sudo ln -sf /usr/include/x86_64-linux-gnu/asm /usr/include/asm

go generate

go generate && go build && sudo ./ebpf-test
```