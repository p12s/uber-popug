import axios from 'axios';
import { setAccount, setToken } from "../reducers/accountReducer";
import { config } from '../config';

export const signUp = (name, username, password) => {
  try {
    axios.post(`${config.AUTH_URL}/sign-up`, JSON.stringify({ 
      name: name, 
      username: username, 
      password: password 
    }), {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then((response) => {
      if (response.status === 200) {
        signIn(username, password)
      } else {
        console.log("возникла ошибка при регистрации")
        alert("возникла ошибка при регистрации")
      }
    })
  } catch (e) {
    alert('возникла ошибка при регистрации')
    console.log(e)
  }
}

export const signIn = (username, password) => {
  try {
    axios.post(`${config.AUTH_URL}/sign-in`, JSON.stringify({
      username: username, 
      password: password 
    }), {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then((response) => {
      if (response.status === 200) {
        setToken(response.data.token)
        localStorage.setItem('token', response.data.token)
        window.location.href = '/tasks';

      } else {
        // TODO вывод ошибок в плашке
        console.log("возникла ошибка при аутентификации")
        alert("возникла ошибка при аутентификации")
      }
    })
  } catch (e) {
    alert('возникла ошибка при аутентификации')
    console.log(e)
  }
}

export const getAccount = () => {
  return async dispatch => {
    try {
      const token = localStorage.getItem('token')

      if (token !== null) {
        const response = await axios.get(`${config.AUTH_URL}/token/`, {
          headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}
        })

        if (response.status === 200) {
          dispatch(setAccount(response.data))
        } else {
          console.log('error getting account data')
        }
      } else {
        window.location.href = '/login';
      }

    } catch (e) {
      console.log('возникла ошибка 211')
      console.log(e)
    }
  }
}