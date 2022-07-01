import './App.css';
import Collection from './Collection';
import Player from './Player';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';

function App() {

  return (
  <Router>
    <Routes>
 <Route exact path="/" element={<Collection />}/> 
 <Route path="/play/:id" element={<Player />}/> 
</Routes>
 </Router>
  );
}

export default App;
