# E-Chat Product Requirements Document (PRD)

## Product Name
E-Chat

---

# 1. Overview

E-Chat is a modern real-time messaging application inspired by WhatsApp, focused on secure communication, social connectivity, and lightweight social sharing.

The platform enables users to:

- Register and create accounts
- Add friends via username search
- Receive friend suggestions based on country and address/location
- Exchange real-time end-to-end encrypted messages
- Create and manage group chats
- Share temporary stories/status updates visible to friends

The application will support both mobile and web clients with a scalable backend architecture.

---

# 2. Goals

## Primary Goals

- Provide fast and secure messaging
- Offer end-to-end encrypted communication
- Build a social graph through friend connections
- Enable lightweight social sharing through stories
- Support scalable group communication

## Secondary Goals

- Enable push notifications
- Support media uploads
- Provide cross-platform support
- Deliver a modern user experience

---

# 3. Target Audience

## Primary Users

- Individuals seeking secure communication
- Friends and family communication groups
- Students and communities
- Small organizations and teams

## Platforms

- Android
- iOS
- Web

---

# 4. Core Features

# 4.1 Authentication & User Accounts

## Description

Users can create accounts and securely authenticate into the platform.

## Functional Requirements

### Registration

Users can register using:

- Email + Password
- Phone Number + OTP (optional future support)

### Login

Users can login using:

- Username/email
- Password

### Profile Setup

Users can:

- Upload profile picture
- Set display name
- Set username
- Set bio/status
- Set country
- Set address/location

### Username Rules

- Must be unique
- Minimum 4 characters
- Only letters, numbers, underscores

Example:

@ebube_dev
@gloria_01
---

# 4.2 Friend System

## Description

Users can search and add friends using usernames.

## Functional Requirements

### Add Friend

- User searches another user by username
- Sends friend request
- Recipient can:
  - Accept
  - Decline

### Friend List

Users can:

- View friends
- Remove friends
- Block users

### Suggested Friends

Suggestions are generated using:

- Same country
- Similar address/location
- Mutual friends (future enhancement)

## Suggestion Algorithm (MVP)

### Priority Scoring

| Factor | Weight |
|---|---|
| Same Country | 50% |
| Same City/Address | 30% |
| Mutual Connections | 20% |

---

# 4.3 Real-Time Messaging

## Description

Users can exchange real-time messages securely.

## Functional Requirements

### Message Types

Support:

- Text messages
- Images
- Videos
- Audio messages
- Documents
- Emojis

### Messaging Features

- Real-time delivery
- Typing indicators
- Read receipts
- Online/offline status
- Message reactions
- Reply to messages
- Forward messages
- Delete for self
- Delete for everyone

### Chat List

Users can:

- Pin chats
- Mute chats
- Archive chats

---

# 4.4 End-to-End Encryption (E2EE)

## Description

Messages are encrypted before leaving the sender device and decrypted only on recipient devices.

## Security Requirements

### Encryption Model

Use:

- Signal Protocol-inspired architecture
- Public/Private key pairs
- Session keys

### Rules

- Server cannot read messages
- Messages encrypted client-side
- Media encrypted before upload

### Stored Data

Server stores:

- Encrypted message payload
- Metadata
- Delivery status

Server DOES NOT store:

- Plain text messages
- Decrypted media

## Security Features

- Device verification
- Key rotation
- Session expiration
- Forward secrecy

---

# 4.5 Group Chats

## Description

Users can create groups for multi-user communication.

## Functional Requirements

### Group Creation

Users can:

- Create group
- Add group image
- Set group name
- Add description

### Group Roles

Roles:

- Owner
- Admin
- Member

### Group Features