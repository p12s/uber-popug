import axios from 'axios';
import { setAccount } from "../reducers/accountReducer";
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
        console.log("success signUp!")
        console.log(response)
        console.log(response.data)
        // TODO вывод плашки успеха?
        // TODO очистка полей

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
        setAccount(response.data.token)
        localStorage.setItem('token', response.data.token)
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