
# TraceCDN
```
    /\_____/\
   /  o   o  \
  ( ==  ^  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
  
```
> A passive subdomain scanner that detects CDN providers.

---

## What it does

- Enumerates subdomains using multiple passive sources
- Detects CDN provider for each subdomain (Cloudflare, Akamai, AWS Cloudfront, Google, Fastly, Sucuri, BunnyCDN)
- No active requests to the target — fully passive

## Sources

- HackerTarget
- RapidDNS
- crt.sh
- urlscan

---

## Installation

Precompiled binaries available for Linux and Android (Termux).

**Linux**
```bash
wget https://github.com/mig-af/CDN-Scanner/raw/refs/heads/main/cdnscanner 
chmod +x cdnscanner
./cdnscanner
```

**Android (Termux)**
```bash
wget https://github.com/mig-af/CDN-Scanner/raw/refs/heads/main/cdnscanner-android 
chmod +x cdnscanner-android
./cdnscanner-android
```

---

## Usage

```
cdnscanner [options] <arguments>

Options:
  --ip <IP>          Scan all CDNs for a given IP
  --domain <DOMAIN>  Scan all CDNs and subdomains for a domain
  --save | -s        Save results to ./subdomains-your-domain.txt
  --help             Show help
```

**Examples**
```bash
./cdnscanner --ip 123.123.123.123
./cdnscanner --domain mydomain.com
./cdnscanner --ip 123.43.45.34 --save
```

---

## Supported CDNs

| CDN | Detection method |
|---|---|
| Cloudflare | IP range |
| Akamai | IP range |
| AWS Cloudfront | IP range |
| Google | IP range |
| Fastly | IP range |
| Sucuri | IP range |
| BunnyCDN | IP range |

---

## Disclaimer

For educational and authorized use only.

---


## Author

- [@Migaf](https://www.github.com/mig-af)

