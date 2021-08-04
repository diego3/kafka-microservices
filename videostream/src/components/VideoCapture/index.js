import React, { useRef, useState, useCallback, useEffect } from 'react'
import Webcam from 'react-webcam'

export default function VideoCapture() {
    const webcamRef = useRef(null)
    let mediaRecorderRef = null
    const [capturing, setCapturing] = useState(false)
    const [recordedChunks, setRecordedChunks] = useState([])

    // https://www.npmjs.com/package/react-webcam
    const videoConstraints = {
        //width: 1280,
        //height: 720,
        //facingMode: "user"
    }
    
    // https://codepen.io/mozmorris/pen/yLYKzyp?editors=0010
    const handleStartCaptureClick = () => {
        setCapturing(true)
        mediaRecorderRef = new MediaRecorder(webcamRef.current.stream, {
          mimeType: "video/webm"
        })

        mediaRecorderRef.ondataavailable = onDataAvailable
        
        mediaRecorderRef.start()
    }

    function onDataAvailable(data) {
        console.log("data", data)
    }

    const handleDataAvailable = useCallback(({ data }) => {
        if (data.size > 0) {
            console.log("data", data)
            //setRecordedChunks((prev) => prev.concat(data));
        }
    },[setRecordedChunks])
    
    const handleStopCaptureClick = useCallback(() => {
        mediaRecorderRef.stop();
        setCapturing(false);
    }, [mediaRecorderRef, webcamRef, setCapturing])

    useEffect(() => {
        console.log("video cam started")
        setTimeout(() => {
            console.log("start capture")
            
        }, 3000)
    }, [])

    function handleUserMedia(e) {
        console.log("userMedia", e)
        mediaRecorderRef = new MediaRecorder(e, {
            mimeType: "video/webm"
        })

        mediaRecorderRef.ondataavailable = onDataAvailable
        
        mediaRecorderRef.start(3000)
    }

    return (
        <>
            <Webcam ref={webcamRef}
                audio={false}
                onUserMedia={handleUserMedia}
            //height={720}
            //width={1280}
            videoConstraints={videoConstraints}
          />
        </>
    )
}