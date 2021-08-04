import { Replay } from "vimond-replay"
import 'vimond-replay/index.css'

// https://github.com/vimond/replay
export default function VideoPlayer1() {
    const bunny = "https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"

    return (
        <>
            <Replay
                source={{
                    streamUrl: 'react-webcam-stream-capture.webm'
                }}
            ></Replay>
        </>
    )
}