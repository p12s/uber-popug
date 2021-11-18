const SET_ACCOUNT = "SET_ACCOUNT"
const SET_TOKEN = "SET_TOKEN"

const defaultState = {
  account: null,
  token: null
}

export default function accountReducer(state = defaultState, action) {
  switch (action.type) {
    case SET_ACCOUNT:
      return {
        ...state,
        token: action.payload
      }
    case SET_TOKEN:
      return {
        ...state,
        token: action.payload
      }
    default:
      return state
  }
}

export const setAccount = (account) => ({type: SET_ACCOUNT, payload: account})
export const setToken = (token) => ({type: SET_TOKEN, payload: token})
