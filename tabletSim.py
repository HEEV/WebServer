import websockets
import argparse
import asyncio
import json


async def interactive(ws):
    while True:
        # Get data to send
        to_send = input("> ")

        # Quit the client
        if to_send == "quit":
            # Send quit over websocket 
            await ws.send("quit")

            # No more input loop
            break

        # Ship that puppy off
        await ws.send(to_send)

        # Get response
        resp = await ws.recv()
        print(f"< {resp}")


async def interactive_init(addr: str, port: int):
    # Connect to websockets on provided address/port
    print(f"Connecting to {addr}:{port}...")

    async with websockets.connect(f"ws://{addr}:{port}/ws") as ws:
        # Send initial packet
        initial = json.dumps(
            {
                "MessageType": "GetNextRunNumber",
                "AndroidID": "0000"
            }
        )
        print(f"> {initial}")
        await ws.send(initial)

        # Get response
        initialResp = await ws.recv()
        print(f"< {initialResp}")

        # Pass off control to actual interactive handler
        await interactive(ws)

    return None

async def sample(addr: str, port: int):
    # Connect to websockets on provided address/port
    print(f"Connecting to {addr}:{port}...")

    # Connect to websocket
    async with websockets.connect(f"ws://{addr}:{port}/ws") as ws:
        # Send initial packet
        initial = json.dumps(
            {
                "MessageType": "GetNextRunNumber",
                "AndroidID": "0000"
            }
        )
        print(f"> {initial}")
        await ws.send(initial)

        # Get response
        initialResp = await ws.recv()
        print(f"< {initialResp}")

    return None

if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Supermileage Tablet Simulator",
    )
    parser.add_argument(
        "-a", "--address", type=str, help="HTTP address to websocket", default="localhost"
    )
    parser.add_argument(
        "-p", "--port", type=int, help="Port to connect on", default=8000
    )
    parser.add_argument(
        "-s", "--sample", action="store_true", default=False, help="Indicates whether to run the sample program"
    )
    
    args = parser.parse_args()

    # Event loop for asyncio operations
    loop = asyncio.get_event_loop()

    if args.sample:
        loop.run_until_complete(sample(args.address, args.port))
    else:
        loop.run_until_complete(interactive_init(args.address, args.port))
