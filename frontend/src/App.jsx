import React, {useState, useEffect} from 'react';
import { BrowserRouter } from 'react-router-dom';
import AppRouter from './components/AppRouter';
import './common.css';
import { AuthContext } from './context';

function App() {
  const [isAuth, setIsAuth] = useState(false)

  useEffect(() => {
    if (localStorage.getItem('token')) {
      setIsAuth(true)
    }
  }, []); 

  return (
    <>
      <AuthContext.Provider value={{ isAuth, setIsAuth }}>
        <BrowserRouter>
          <AppRouter/>
        </BrowserRouter>
      </AuthContext.Provider>
    </>
  );
}

export default App;
