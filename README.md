# Beamsend

## Idea

Creating a **file-sharing system** that's like a universal AirDrop.

- It will be fast and lightweight.
- It will be made in Go and vanilla javascript.
- It will have very minimal UI.

The user can upload any type of files up to a certain size, then a link will be generated for downloading the uploaded files.
Once the files are downloaded, the server will delete the files, but will store:

1. Which device uploaded the files
2. The IP address
3. The name of the files uploaded.

The file-sharing system will also support direct connection without link generation.

## Improvements

### Features

- Implement WebRTC for for peer-to-peer file tranfer; eliminating the need for a server in *direct connect* scenario.

### Security

- HTTPS connections.
- Implement strong encryption for uploaded files both at rest and in transit.
- Consider using short-lived, unique download links with expiration times to further enhance security.
- You'll need a strategy for user authentication (even if basic) to manage access and prevent unauthorized uploads/downloads.

## RoadMap

1. **Core Functionality**
    - [x] Develop a file-upload with temporary storage using Go.
    - [ ] Implement download link generation with basic security measures.
    - [ ] Build the minimal UI using vanilla JavaScript.
2. **Direct Connection**
    - [ ] Research WebRTC integration for peer-to-peer file transfer.
    - [ ] Develop logic for device discovery & secure connection establishment.
3. **Advance Features**
    - [ ] PWA version

## Closing

### Naming

\[image here]
