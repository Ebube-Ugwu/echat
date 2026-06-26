# E-Chat Implementation Plan

**Pace:** 40 minutes per working day
**Estimated duration:** ~52 working days (~10.5 weeks)
**Approach:** Each feature phase combines backend + frontend — feature is fully demoable before moving on

---

## Phase 1: Foundation & Auth (Days 1–8)

### Day 1 — Project Scaffolding
- **Backend:** Fix broken `V1__create_users_table.sql` migration; add PostgreSQL JDBC driver; configure `application.properties` (DB, port); verify Flyway + Spring Boot start
- **Frontend:** Scaffold React + TypeScript project with Vite; install Tailwind CSS, react-router, axios, zustand; create folder structure

### Day 2 — User Entity & Frontend Setup
- **Backend:** Create `User` JPA entity (id, username, email, password, displayName, bio, country, address, profilePic); create `UserRepository`; write `V2__add_user_details.sql` migration
- **Frontend:** Set up routing (Login, Register, Home, Profile pages); create `AuthContext` + `AuthService` with axios interceptor; wire up global auth state in zustand

### Day 3 — Registration (Backend)
- Create `AuthController` with `POST /api/auth/register`
- Create `AuthService` with validation, password hashing (BCrypt)
- Create DTOs and error responses for duplicate username/email

### Day 4 — Registration (Frontend)
- Build Register page (username, email, password, displayName form)
- Connect to backend register endpoint
- Show validation errors inline
- On success, store JWT and redirect to profile setup

### Day 5 — Login & JWT (Backend)
- Add Spring Security; configure `SecurityConfig` (permit auth endpoints)
- Implement `JwtTokenProvider` (generate/validate)
- Create `POST /api/auth/login` returning JWT
- Create `JwtAuthenticationFilter`

### Day 6 — Login (Frontend) & Auth Guard
- Build Login page (username/email + password form)
- Connect to backend login endpoint
- Create `ProtectedRoute` component (redirects to /login if unauthenticated)
- Wire token expiry → auto-logout

### Day 7 — Profile Endpoints (Backend)
- Create `GET /api/auth/me` (current user from token)
- Create `PUT /api/users/profile` (update displayName, bio, country, address)
- Create avatar upload endpoint `POST /api/users/avatar`
- `UserProfileResponse` DTO (no password hash)

### Day 8 — Profile Pages (Frontend)
- Build Profile Setup page (first-login flow: avatar, displayName, bio, country, address)
- Build Profile Edit page (edit all fields, preview avatar)
- Connect all profile endpoints
- **Verify:** Full auth flow works end-to-end (register → login → profile → edit)

✅ **Milestone 1: Auth is fully working with UI**

---

## Phase 2: Friend System (Days 9–14)

### Day 9 — Friend Request Backend
- Create `FriendRequest` JPA entity (sender, recipient, status: PENDING/ACCEPTED/DECLINED/BLOCKED)
- Write `V3__create_friend_requests_table.sql`
- Create `FriendRequestRepository` + `FriendService`

### Day 10 — Friend APIs (Backend)
- `POST /api/friends/request/{username}` — send request
- `PUT /api/friends/request/{requestId}/accept` — accept
- `PUT /api/friends/request/{requestId}/decline` — decline
- `DELETE /api/friends/{friendId}` — remove friend
- `POST /api/friends/block/{userId}` / `DELETE /api/friends/unblock/{userId}`

### Day 11 — Friend List & Search APIs
- `GET /api/friends` — list accepted friends
- `GET /api/friends/requests` — list pending (incoming + outgoing)
- `GET /api/users/search?q={username}` — search users
- `GET /api/friends/suggestions` — friend suggestions (country 50%, address 30%)

### Day 12 — Friend System UI (Frontend)
- Build Friends page with tabs: All Friends | Requests | Search
- Search users with debounced input, display results with "Add Friend" button
- Accept/decline request buttons on incoming requests
- Remove friend confirmation dialog

### Day 13 — Friend Suggestions UI & Blocking
- Build Friend Suggestions list on Friends page (with score badge)
- Block/unblock UI (from friend's profile context menu)
- Show blocked users in settings section

### Day 14 — Polish & Integration Test
- End-to-end test: search user → send request → accept → appears in friend list → verify suggestion algorithm
- Handle edge cases: self-request, duplicate request, re-request after decline
- **Verify:** Full friend flow works end-to-end

✅ **Milestone 2: Friend system fully working with UI**

---

## Phase 3: Real-Time Messaging Core (Days 15–21)

### Day 15 — WebSocket & Chat Entities (Backend)
- Add WebSocket dependency; configure STOMP with `/queue` and `/topic` prefixes
- Create `Chat`, `Message`, `ChatParticipant` JPA entities
- Write `V4__create_chat_and_message_tables.sql` migration
- Create repositories for all three

### Day 16 — Chat Service & REST APIs (Backend)
- `ChatService.createPrivateChat(user1, user2)` — find-or-create
- `GET /api/chats` — list user's chats with last message + unread count
- `GET /api/chats/{chatId}/messages?page=X&size=Y` — paginated history

### Day 17 — WebSocket Messaging (Backend)
- STOMP `@MessageMapping /chat.send` — persist message, deliver via `/queue/messages/{userId}`
- WS authentication from STOMP session headers
- `GET /api/users/{userId}/status` — online/offline status
- Track connections in memory for online status

### Day 18 — Chat List UI (Frontend)
- Build Chat List sidebar (avatar, name, last message preview, timestamp, unread badge)
- Search/filter chats by name
- Sort by most recent activity
- Wire up to `GET /api/chats` endpoint

### Day 19 — Chat Screen UI (Frontend)
- Build Chat page with message bubbles (sent vs received styling)
- Message input with send button
- Connect WebSocket on chat open
- Handle incoming messages in real-time via zustand store
- Auto-scroll to bottom

### Day 20 — Online Status & Typing (Frontend)
- Show online/offline indicator next to chat partner name
- `@MessageMapping /chat.typing` — send typing indicator
- Display "typing..." indicator when partner is typing
- Wire to backend typing STOMP messages

### Day 21 — Integrated Chat Polish
- Handle WebSocket reconnection on network loss
- Fix any edge cases: empty chat, sending to offline user, etc.
- **Verify:** Two users can register, become friends, open a chat, and exchange messages in real-time

✅ **Milestone 3: Real-time 1-on-1 messaging fully working**

---

## Phase 4: Advanced Message Features (Days 22–26)

### Day 22 — Reactions & Reply (Backend)
- Create `MessageReaction` entity; write `V5__create_message_reactions_table.sql`
- `POST /api/chats/{chatId}/messages/{msgId}/react` + `DELETE` to remove
- Add `replyToId` field to Message entity migration
- Populate replyToId on message send if replying

### Day 23 — Reactions & Reply UI (Frontend)
- Emoji picker on hover/long-press for reactions
- Display reactions below messages (inline emoji + count)
- Tap to reply: quoted message preview above input bar
- Send reply message with quoted context

### Day 24 — Read Receipts (Backend + Frontend)
- **Backend:** Create `MessageReadReceipt` entity; write `V6__create_read_receipts_table.sql`; `POST /api/chats/{chatId}/messages/{msgId}/read`
- **Frontend:** Show double-check (✓✓) indicator on sent messages; display "Read by {name}" on info tap; mark messages as read when chat is opened

### Day 25 — Forward & Delete (Backend + Frontend)
- **Backend:** `POST /api/chats/{chatId}/messages/{msgId}/forward` — copy to target chat; `DELETE` — delete for self (soft delete); delete for everyone (sender, time-limited)
- **Frontend:** Forward message UI (select target chat from list); delete context menu with "Delete for me" / "Delete for everyone"

### Day 26 — Pin, Mute, Archive (Backend + Frontend)
- **Backend:** Add `isPinned`, `isMuted`, `isArchived` fields to `ChatParticipant`; write `V7` migration; `PUT /api/chats/{chatId}/pin|mute|archive`
- **Frontend:** Chat context menu with pin/mute/archive toggles; show pin indicator; filter archived chats; muted chat indicator
- **Verify:** Full messaging feature set working

✅ **Milestone 4: All 1-on-1 messaging features complete**

---

## Phase 5: End-to-End Encryption (Days 27–30)

### Day 27 — Key Management Backend
- Create `UserKeyBundle` entity (identityKeyPub, signedPreKeyPub, preKeySig, oneTimePreKeys[])
- Write `V8__create_user_key_bundle_table.sql`
- `POST /api/keys/upload` — user uploads key bundle
- `GET /api/keys/{userId}` — get user's public key bundle

### Day 28 — Encryption Service (Backend)
- Create `EncryptionService` using Java Cryptography Architecture
- Implement AES-256-GCM for message content encryption
- Per-conversation key exchange using RSA/OAEP
- Encrypt message content before DB persist; decrypt on read

### Day 29 — Integrate Encryption into Messaging
- On private chat creation, establish shared session key
- Encrypt message content before storage (server sees only ciphertext)
- Decrypt on WebSocket delivery to recipient
- Add encryption metadata to Message entity (keyId, iv)
- Verify in logs that plaintext is never logged

### Day 30 — Encryption Integration (Frontend)
- Generate and upload key bundle on first login (after registration)
- Fetch partner's public key when opening a chat
- For MVP: indicate encryption status with lock icon per message
- Show "Messages are end-to-end encrypted" banner in chat
- **Verify:** Messages in database are ciphertext; UI shows encryption indicator

✅ **Milestone 5: End-to-end encryption operational**

---

## Phase 6: Group Chats (Days 31–36)

### Day 31 — Group Backend Basics
- `POST /api/groups` — create group (name, description, image, initial members)
- `GET /api/groups` — list user's groups
- `PUT /api/groups/{id}` — update group info (owner/admin)
- Write `V9__create_groups_table.sql` migration (group entities use existing Chat + ChatParticipant with group-specific fields)

### Day 32 — Group Membership Backend
- `POST /api/groups/{id}/members` — add members (admin only)
- `DELETE /api/groups/{id}/members/{userId}` — remove member
- `PUT /api/groups/{id}/members/{userId}/role` — promote/demote (owner only)
- `PUT /api/groups/{id}/leave` — leave group (transfer ownership if owner)

### Day 33 — Group Messaging Backend
- Extend WebSocket messaging to groups: `/topic/groups/{groupId}`
- Fan-out delivery to all online members
- Group read receipts (track per-member)
- `GET /api/groups/{id}/messages` — paginated history

### Day 34 — Group UI (Frontend)
- Build Create Group form (name, description, select members from friend list)
- Group info page (members list with roles, avatar, leave button)
- Show group name + member count in chat header

### Day 35 — Group Messaging UI & Admin Controls
- Group chat in existing Chat screen (show sender name per message)
- Admin controls: add members dialog, manage roles dropdown
- Leave group with confirmation

### Day 36 — Group Polish
- Handle edge cases: group with single member, re-adding removed member, group avatar upload
- **Verify:** Create group → add members → group messaging works in real-time

✅ **Milestone 6: Group chats fully working**

---

## Phase 7: Stories / Status (Days 37–41)

### Day 37 — Story Backend Core
- Create `Story` entity (userId, mediaUrl, content, type, expiresAt = createdAt + 24h)
- Write `V10__create_stories_table.sql`
- `StoryRepository` with `findAllByUserIdInAndExpiresAtAfter` query
- `POST /api/stories` — create story

### Day 38 — Story APIs & Views Backend
- `GET /api/stories` — get non-expired stories from friends
- `DELETE /api/stories/{id}` — delete own story
- Create `StoryView` entity; write `V11__create_story_views_table.sql`
- `POST /api/stories/{id}/view` — mark as viewed
- `GET /api/stories/{id}/viewers` — return viewer list + count

### Day 39 — Story UI (Frontend)
- Build Stories section: horizontal scroll of circular story avatars at top of home page
- Ring indicator: seen (gray) vs unseen (color)
- "Add Story" button (camera icon on user's own avatar)

### Day 40 — Story Viewer & Creator (Frontend)
- Build Story Viewer: full-screen overlay, tap left/right to navigate, auto-advance, 24h timer countdown
- Build Create Story modal: text overlay, optional image upload
- Show view count on own stories

### Day 41 — Story Polish
- Auto-refresh stories when 24h expires
- Empty state when no friends have stories
- **Verify:** Create story → appears in friends' story feeds → view tracking works

✅ **Milestone 7: Stories feature fully working**

---

## Phase 8: Media Uploads & Push Notifications (Days 42–45)

### Day 42 — File Upload Backend
- Create `FileStorageService` (local filesystem; configurable path)
- `POST /api/upload` — upload file, return URL
- Restrict file types (jpg/png/mp4/mp3/pdf); enforce size limits
- Generate unique filenames; encrypt files before storage

### Day 43 — Media Messages Backend
- Update message send flow to support media URLs
- Generate thumbnails for images
- Wire media into E2EE (encrypt media file before storing)
- Clean up files when messages are deleted

### Day 44 — Push Notifications Backend
- Create `DeviceToken` entity; write `V12__create_device_tokens_table.sql`
- `POST /api/devices/register` — register FCM/APNs token
- Create `PushNotificationService`
- Trigger push on: new message (offline recipient), friend request, story from friend

### Day 45 — Media & Push Frontend Integration
- Media picker in message input (image, video, document attach buttons)
- Display media messages in chat (image preview, video player, document download)
- Image preview lightbox on tap
- Handle notification click → deep link to chat
- Register device token on login

✅ **Milestone 8: Media sharing & push notifications working**

---

## Phase 9: Polish & Deployment (Days 46–52)

### Day 46 — Error Handling & UX
- **Backend:** Global `@ControllerAdvice` exception handler; standardized error response format; rate limiting on auth endpoints
- **Frontend:** Global error boundary; toast notifications for API errors; loading skeletons; offline banner when disconnected

### Day 47 — Backend Testing
- Unit tests for `AuthService` (register, login, validation)
- Unit tests for `FriendSuggestionService` (scoring algorithm)
- Unit tests for `EncryptionService` (encrypt/decrypt round-trip)
- Integration tests for key REST endpoints

### Day 48 — Frontend Testing
- Component tests for Login, Register, Chat, Friend pages
- Test auth flow (login → token → protected routes)
- Test message send/receive WebSocket flow

### Day 49 — Docker & Deployment Config
- Multi-stage `Dockerfile` for backend (Jar build → runtime)
- `Dockerfile` for frontend (build → nginx serve)
- `docker-compose.yml` (backend + frontend + PostgreSQL)
- Environment variable configuration (.env.example files)

### Day 50 — Security & Performance Audit
- Verify all endpoints authenticated (except auth routes)
- Check password hashing strength (BCrypt cost factor)
- Verify E2EE: DB contains only ciphertext; no plaintext in logs
- Add `@PreAuthorize` checks where missing
- Quick perf check: pagination limits, N+1 query prevention

### Day 51 — Full User Journey Walkthrough
- Register → set up profile → search/add friends → 1-on-1 chat → send reaction/reply → forward → delete
- Create group → add members → group chat
- Create story → view as friend
- Upload media in message
- Test offline behavior and reconnection

### Day 52 — Final Bug Fixes & Documentation
- Fix any remaining issues from walkthrough
- Update README with setup instructions
- Record known limitations and future improvements
- Final commit

---

## Summary

| Phase | Days | Feature | Demoable After |
|-------|------|---------|----------------|
| 1 | 1–8 | Auth & Profiles | Day 8 ✓ |
| 2 | 9–14 | Friend System | Day 14 ✓ |
| 3 | 15–21 | Real-Time Messaging | Day 21 ✓ |
| 4 | 22–26 | Message Features | Day 26 ✓ |
| 5 | 27–30 | E2EE | Day 30 ✓ |
| 6 | 31–36 | Group Chats | Day 36 ✓ |
| 7 | 37–41 | Stories | Day 41 ✓ |
| 8 | 42–45 | Media & Push | Day 45 ✓ |
| 9 | 46–52 | Polish & Deploy | Day 52 ✓ |

**Total: 52 working days = ~10.5 weeks @ 40 min/day**
