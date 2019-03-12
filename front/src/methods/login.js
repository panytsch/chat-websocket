import axios from "axios"

export const host = '127.0.0.1:8000/api';

export const login = (name, pass) => dispatch => {
    if (!name || !pass) return null;

    axios.get(`/api/auth`, {
        headers: {
            name: name,
            password: pass,
        }
    }).then(({data}) => {
        console.log(data);
        if (!data.token) return;
        dispatch({
            type: 'AUTH',
            payload: data.token
        })
    })
};