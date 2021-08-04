import './App.css';

import CaptureTutorial from './components/CaptureTutorial'
import VideoCapture from './components/VideoCapture'
import VideoPlayer1 from './components/VideoPlayer1'

function App() {

  
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Video streaming
        </p>
        
        <div className="container">
          <div className="my-webcam">
              <CaptureTutorial></CaptureTutorial>
          </div>
          <div className="my-player">
              <VideoPlayer1></VideoPlayer1>
          </div>
        </div>
      </header>
    </div>
  );
}

export default App;
