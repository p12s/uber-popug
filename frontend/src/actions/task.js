import axios from 'axios';
import { setTasks } from "../reducers/taskReducer";
import { config } from '../config';

export const getAllTask = () => {
  return async dispatch => {
    try {
      const token = localStorage.getItem('token')

      if (token !== null) {
        const response = await axios.get(`${config.TASK_URL}/task/`, {
          headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}
        })

        if (response.status === 200) {
          if (!response.data) {
            dispatch(setTasks([]))
          } else {
            dispatch(setTasks(response.data))
          }
        } else {
          window.location.href = '/login';
        }
      } else {
        window.location.href = '/login';
      }

    } catch (e) {
      console.log('возникла ошибка 2222')
      console.log(e)
    }
  }
}

export const setTaskCompleted = () => {
  return async dispatch => {
    try {
      const token = localStorage.getItem('token')

      if (token !== null) {
        // const response = await axios.get(`${config.TASK_URL}/task/`, {
        //   headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}
        // })

        // if (response.status === 200) {
        //   if (!response.data) {
        //     dispatch(setTasks([]))
        //   } else {
        //     dispatch(setTasks(response.data))
        //   }
        // } else {
        //   window.location.href = '/login';
        // }
      } else {
        window.location.href = '/login';
      }

    } catch (e) {
      console.log('возникла ошибка setTaskCompleted')
      console.log(e)
    }
  }
}
