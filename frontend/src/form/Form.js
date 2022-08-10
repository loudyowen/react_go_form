import React, {useState} from "react";
import {TextField, Paper, Button} from '@mui/material';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import useStyles from './styles'
import {useDispatch} from 'react-redux'
import { createForm } from "../actions/formAction";
const Form = () =>{
    const classes = useStyles();
    const dispatch = useDispatch();
    const [startValue, setStartValue] = React.useState(new Date());
    const [endValue, setEndValue] = React.useState(new Date());
    const [formData, setFormData] = useState({
        nama: '',
        reason: '',
    })

    const handleSubmit = (e) => {
        e.preventDefault();
        dispatch(createForm({...formData,startValue,endValue}))
    }
    return(
        <Paper className={`${classes.paper} ${classes.root}`} elevation={6}>
             <h1>
                Form Overtime Request
            </h1>
            <form className={classes.form} autoComplete="off" noValidate onSubmit={handleSubmit}>
                <TextField name='name' id="outlined-basic" label="Nama" variant="outlined" fullWidth value={formData.nama} onChange={(e)=>setFormData({...formData, nama: e.target.value})}/>
                <TextField name='reason' id="outlined-textarea" multiline label="Alasan" variant="outlined" fullWidth rows={4} value={formData.reason} onChange={(e)=>setFormData({...formData, reason: e.target.value})}/>
            

                <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DateTimePicker
                    renderInput={(props) => <TextField {...props} />}
                    label="StartTime"
                    value={startValue}
                    ampm={false}
                    disableFuture={true}
                    onChange={(newStartValue) => {
                    setStartValue(newStartValue);
                    }}
                />
                <DateTimePicker
                    renderInput={(props) => <TextField {...props} />}
                    label="EndTime"
                    value={endValue}
                    ampm={false}
                    disableFuture={true}
                    onChange={(newEndValue) => {
                        setEndValue(newEndValue);
                    }}
                />
                </LocalizationProvider>
                <Button variant="contained" color="primary" size="large" type="submit">Submit</Button>
            </form>
        </Paper>

    )
}

export default Form