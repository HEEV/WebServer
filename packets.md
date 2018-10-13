# Packets

All communication will be done via JSON, with the packet structure as described
below.

## Packet Definitions

All packets will follow the same basic form, with the same required
identification packets whether the packet is from a client or the server.

### Identification

All packets being sent to/from the server must have these fields. If they are
not included the packet is invalid and will be discarded and ignored.

- `messageType` - A string representing what data this message is expected to contain
- `androidId` - A string representing the ID of the android tablet this data is going to/coming from.

This should be unique between all clients. 

If a client is not actually an android tablet then this value can be anything that will be unique.