import React from "react";
import './App.css';
import Form from './form/Form';
import {BrowserRouter, Routes, Route} from "react-router-dom"

function App() {
  return (
    <BrowserRouter>
    <Routes>
      <Route path="/" exact element={<Form />} />
    </Routes>
    </BrowserRouter>
  );
}

export default App;
