# Real-Time Crowd-Aware Navigation with Offline-First Conflict Resolution and Temporal State Replay

## Overview

This project is an offline-first, crowd-aware navigation system designed for high-density urban environments where reliable connectivity cannot always be guaranteed.

The system combines a Flutter mobile application with a Go backend and Firebase Realtime Database to handle real-time location synchronization, offline location recording, temporal event replay, deterministic state reconstruction, crowd-aware routing, conflict resolution, and route auditability.

The core challenge is that location events may be:

- Generated while the device is offline
- Received after a network connection is restored
- Received out of chronological order
- Received multiple times because of synchronization retries
- Conflicting with previously received location updates

The system therefore does not rely only on event arrival order. Instead, location events contain deterministic metadata such as timestamps, sequence numbers, event IDs, and device IDs so that the user's state can be reconstructed and replayed consistently.

---

# Problem Statement

In a high-density urban environment, real-time crowd information is important for safe navigation.

Users may frequently enter low-connectivity areas such as:

- Subways
- Tunnels
- Underground stations
- Areas with unstable mobile connectivity

During these periods, location updates cannot always be sent immediately to the backend.

When connectivity returns, multiple events may arrive together, potentially:

- Out of order
- Late
- Duplicated
- Conflicting with newer events

A conventional "last event received wins" approach can produce an incorrect user state.

This project solves the problem using an offline-first event-based architecture.

The system:

1. Stores location events locally while offline.
2. Synchronizes pending events when connectivity returns.
3. Processes events using deterministic temporal information.
4. Reconstructs user state through event replay.
5. Resolves duplicates and conflicting updates deterministically.
6. Evaluates live crowd-density zones.
7. Avoids unsafe Red zones when computing routes.
8. Recomputes routes when crowd conditions require it.
9. Generates an audit record for every route recomputation.
10. Allows historical decisions to be reproduced through event replay.

---

# System Architecture

```text
                         USER
                          |
                          v
                 Flutter Mobile App
                          |
                          v
                Location Event Recorder
                          |
                          v
                     SQLite
                          |
              +-----------+-----------+
              |                       |
           OFFLINE                  ONLINE
              |                       |
              |                       v
              |              Firebase Realtime DB
              |                       |
              +-----------+-----------+
                          |
                          v
                    Event Stream
                          |
                          v
                Temporal Event Processing
                          |
                          v
                 Event Replay Engine
                          |
                          v
                State Reconstruction
                          |
                          v
                 Crowd Zone Analysis
                          |
                 +--------+--------+
                 |        |        |
               GREEN   YELLOW     RED
                 |        |        |
                 |        |     Avoid
                 |        |        |
                 +--------+--------+
                          |
                          v
                  Routing Engine
                          |
                          v
                Route Recomputation
                          |
                 +--------+--------+
                 |                 |
                 v                 v
           Route Decision      Audit Record