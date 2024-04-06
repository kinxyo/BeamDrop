> [!IMPORTANT]
> Soon to be deployed on AWS & support WebRTC peer-to-peer connection.

# Beamsend

## About

Creating a **file-sharing system** that's like a universal AirDrop.

**Features**:

- Fast and lightweight, due to being developed using Go and vanilla javascript.
- Minimal UI for extremely ease-of-use.
- Supports any type of files, upto the size of 10GB.
- Drag and drop
- Privacy support: No files are stored on server & No logs created.
- Option to self-host

## Improvements

### Features

- [ ] Implement WebRTC for for peer-to-peer file tranfer; eliminating the need for storing in server.

### Security

- [ ] Add expiration time to the generated links.
- [ ] Implement strong encryption for uploaded files both at rest and in transit.
- [ ] Implement strategy for user authentication (even if basic) to manage access and prevent unauthorized uploads/downloads.
- [ ] HTTPS connections.

## RoadMap

1. **Core Functionality**
    - [x] Develop a file-upload with temporary storage using Go.
    - [x] Implement download link generation
    - [x] Build the minimal UI using vanilla JavaScript.
    - [ ] Improve basic security measures.
2. **Direct Connection**
    - [ ] Research WebRTC integration for peer-to-peer file transfer.
    - [ ] Develop logic for device discovery & secure connection establishment.
3. **Production**
    - [ ] PWA version
    - [ ] AWS deployment
