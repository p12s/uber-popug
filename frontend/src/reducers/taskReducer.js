const SET_TASKS = "SET_TASKS"

const defaultState = {
  tasks: []
}

export default function taskReducer(state = defaultState, action) {
  switch (action.type) {
    case SET_TASKS:
      return {
        ...state,
        tasks: action.payload
      }
    default:
      return state
  }
}

export const setTasks = (tasks) => ({type: SET_TASKS, payload: tasks})
