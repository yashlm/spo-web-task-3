import React, { useState } from 'react';
import axios from 'axios';
import './App.css'

const SignupForm = () => {
  const [userId, setUserId] = useState('');
  const [password, setPassword] = useState('');

  const handleSignup = async (e) => {
    e.preventDefault();

    try {
      // Send signup request to the backend
      await axios.post('http://localhost:8080/signup', { userId, password });
      console.log('Signup successful');
      // Reset the form
      setUserId('');
      setPassword('');
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <form onSubmit={handleSignup}>
      <h1>Signup</h1>
        <div>
          <label>User ID:</label>
          <input type="text" value={userId} onChange={(e) => setUserId(e.target.value)} required />
        </div>
        <div>
          <label>Password:</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        <button type="submit">Signup</button>
      </form>
    </div>
  );
};

export default SignupForm;
