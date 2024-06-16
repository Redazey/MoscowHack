import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './pages/login_page';
import Register from './pages/registration_page';
import MainPage from './pages/main_page';

function App() {
  return (
    <Router>
        <Routes>
            <Route path="/" element={<MainPage />}/>
            <Route path="Login" element={<Login />}/>
            <Route path="Registration" element={<Register />}/> 
        </Routes>
    </Router>
  )
}

export default App
