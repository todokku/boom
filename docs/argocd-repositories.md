# Repositories

For a repository there are two types, with ssh-connection where an url and a certifacte have to be provided and an https-connection where an URL, username and password have to be provided.

| Parameter                          | Description                                                                     | Default                           |
| ---------------------------------- | ------------------------------------------------------------------------------- | --------------------------------- |
| `url`                              | Prefix where the credential should be used (starting "git@" or "https://" )                   |                                   |
| `usernameSecret`                   | Attributes for username in a secret                                             |                                   |
| `usernameSecret.name`              | Name of the secret                                                              |                                   |
| `usernameSecret.key`               | Key in the secret which contains the username                                   |                                   |
| `passwordSecret`                   | Attributes for username in a secret                                             |                                   |
| `passwordSecret.name`              | Name of the secret                                                              |                                   |
| `passwordSecret.key`               | Key in the secret which contains the password                                   |                                   |
| `certificateSecret`                | Attributes for username in a secret                                             |                                   |
| `certificateSecret.name`           | Name of the secret                                                              |                                   |
| `certificateSecret.key`            | Key in the secret which contains the certificate                                |                                   |