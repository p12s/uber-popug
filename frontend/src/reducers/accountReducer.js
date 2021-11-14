const SET_ACCOUNT = "SET_ACCOUNT"

const defaultState = {
  accountId: 0, // TODO убрать если не использую потом
  token: null
}

export default function accountReducer(state = defaultState, action) {
  switch (action.type) {
    case SET_ACCOUNT:
      return {
        ...state,
        token: action.payload
      }
    default:
      return state
  }
}

export const setAccount = (account) => ({type: SET_ACCOUNT, payload: account})