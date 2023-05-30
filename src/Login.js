import React, { useState } from 'react';
import axios from 'axios';
import './App.css'
import SignupForm from './SignUp';

const LoginForm = () => {
  const [userId, setUserId] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      // Send login request to the backend
      await axios.post('http://localhost:8080/login', { userId, password });
      console.log('Login successful');
      // Reset the form
      setUserId('');
      setPassword('');
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className='forms'>
      <form onSubmit={handleLogin}>
      <h1>Login</h1>
        <div>
          <label>User ID:</label>
          <input type="text" value={userId} onChange={(e) => setUserId(e.target.value)} required />
        </div>
        <div>
          <label>Password:</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        <button type="submit">Login</button>
      </form>
      <SignupForm/>
    </div>
  );
};

export default LoginForm;
