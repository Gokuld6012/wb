# WB

Simple website blacklisting using hosts file

## To block website

```bash
sudo wb b twitter.com

sudo wb b facebook.com
```

## To unblock website

```bash
sudo wb ub twitter.com

sudo wb ub facebook.com
```

## TODO

- [ ] Handle url validation errors.
- [ ] Hanlde non existent websites in /etc/hosts file.
