import React, {useState} from "react";
import {TextField, Paper, Button, Container} from '@mui/material';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import useStyles from './styles'
import {useDispatch} from 'react-redux'
import { createForm } from "../actions/formAction";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import timezone from "dayjs/plugin/timezone";

dayjs.extend(utc)
dayjs.extend(timezone)
dayjs.tz.setDefault("America/New_York")


const Form = () =>{
    const classes = useStyles();
    const dispatch = useDispatch();
    const [startDate, setStartDate] = useState(new Date());
    const [endDate, setEndDate] = useState(new Date());
    const [formData, setFormData] = useState({
        nama: '',
        reason: '',
        startDate: '',
        endDate: ''
    })

    const handleSubmit = (e) => {
        e.preventDefault();
        formData.startDate = startDateFormat
        formData.endDate = endDateFormat
        dispatch(createForm({...formData}))
    }

    const startDateFormat = dayjs(startDate).format("YYYY-MM-DD HH:mm")
    const endDateFormat = dayjs(endDate).format("YYYY-MM-DD HH:mm")
    return(
        <Container>
            <Paper className={`${classes.paper} ${classes.root}`} elevation={6} >
                <h1>
                    Form Overtime Request
                </h1>
                <form className={classes.form} autoComplete="off" noValidate onSubmit={handleSubmit}>
                    <TextField name='name' id="outlined-basic" label="Nama" variant="outlined" fullWidth value={formData.nama} onChange={(e)=>setFormData({...formData, nama: e.target.value})}/>
                    <TextField name='reason' id="outlined-textarea" multiline label="Alasan" variant="outlined" fullWidth rows={4} value={formData.reason} onChange={(e)=>setFormData({...formData, reason: e.target.value})}/>
                

                    <LocalizationProvider dateAdapter={AdapterDateFns} dateLibInstance={dayjs.tz}>
                        <DateTimePicker
                            renderInput={(props) => <TextField {...props} />}
                            label="StartTime"
                            value={startDate}
                            ampm={false}
                            disableFuture={true}
                            inputFormat="E dd MM yyyy hh:mm"
                            onChange={(newStartDate) => {
                                setStartDate(newStartDate);
                            }}

                        />
                        <DateTimePicker
                            renderInput={(props)=><TextField {...props}/>}
                            label="EndTime"
                            value={endDate}
                            ampm={false}
                            disableFuture={true}
                            inputFormat="E dd MM yyyy hh:mm"
                            onChange={(newEndDate) => {
                                setEndDate(newEndDate);
                            }}
                        />
                    </LocalizationProvider>
                    <Button variant="contained" color="primary" size="large" type="submit">Submit</Button>
                </form>
            </Paper>
        </Container>


    )
}

export default Form