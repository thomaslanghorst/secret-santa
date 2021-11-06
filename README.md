# secret-santa
This is a small project I wrote to help my family picking secret santas without having a person beeing involved. This way noone is left out of having fun.

The code first reads in the `contacts.csv` file and picks random matches. Afterwards it uses [this whatsapp libary](https://github.com/Rhymen/go-whatsapp) to send massages to each person.

# Setup
1. Before running the code you need to prepare the `contacts.csv` file. In there you need to add lines for each participant of the your secret santa game. Add the name of the contact, followed by a semicolon and his or her phone number in the form `0<international_calling_code>number`. An example looks like this:
    ```
    Thomas;049123456789
    ```
2. Additionally, you need to set the `selfContactName` variable to the same name you gave yourself in the `contacts.csv` file. This way the message you send to yourself will not be deleted.
3. You might also need to set the `whatsappVersionMajor`,  `whatsappVersionMinor` and `whatsappVersionPatch`variables in the `main.go` file


# Running the code
To execute the code simply call
```
go mod *.go
```
and scan the QR code printed on the terminal with your phone.
