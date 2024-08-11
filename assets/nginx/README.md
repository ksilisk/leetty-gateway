# How to prepare reverse-proxy by nginx

## Generate certificates
    $ openssl req -newkey rsa:2048 -sha256 -nodes -keyout YOURPRIVATE.key -x509 -days 365 -out YOURPUBLIC.pem -subj "/C=US/ST=New York/L=Brooklyn/O=Example Brooklyn Company/CN=YOURDOMAIN.EXAMPLE"
`YOURPRIVATE` - is name of your private key file.

`YOURPUBLIC` - is name of your public key file.

`YOURDOMAIN.EXAMPLE` - is IMPORTANT. Should be your server IP-address or URL. Also, should equal to `server_name` in your `nginx.conf`. 
## Enable reverse proxy
1. Fill data to [nginx config sample](./nginx.conf_sample)
2. Add filled config to `/etc/nginx/sites-available`
3. Create symlink

        $ sudo ln -s /etc/nginx/sites-available/YOUR_FILE_NAME.conf /etc/nginx/sites-enabled/YOUR_FILE_NAME.conf
4. Reload nginx to enable reverse-proxy

        $ sudo systemctl reload nginx