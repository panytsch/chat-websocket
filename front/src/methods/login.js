import axios from "axios"

export const host = '127.0.0.1:8000/api';

export const login = dispatch => (name, pass) => {
    if (!name || !pass) return null;

    axios.get(`/api/auth`, {
        headers: {
            name: name,
            password: pass,
        }
    }).then(res => {
        console.log(res);
        dispatch({
            type: 'AUTH',
        })
    })
};