import React, { useState } from 'react';
import sayyes from './images/sayyes.jpg';
import './Form.css';

const Form = () => {
  const [first, setFirst] = useState('');
  const [last, setLast] = useState('');
  const [email, setEmail] = useState('');
  const [phone ,setPhone] = useState('');
  const [company, setCompany] = useState('');
  const [position, setPosition] = useState('');
  const [submitted, setSubmitted] = useState(false);


  const handleFirstChange = (e) => {
    setFirst(e.target.value);
  };

  const handleLastChange = (e) => {
    setLast(e.target.value);
  };

  const handleEmailChange = (e) => {
    setEmail(e.target.value);
  };

  const handlePhoneChange = (e) => {
    setPhone(e.target.value);
  };

  const handleCompanyChange = (e) => {
    setCompany(e.target.value);
  };

  const handlePositionChange = (e) => {
    setPosition(e.target.value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    // Perform form submission logic here
    console.log('First:', first);
    console.log('Last:', last)
    console.log('Email:', email);
    console.log('Phone:', phone);
    console.log('Company:', company);
    console.log('Position:', position);

    setSubmitted(true);

    const response = await fetch("https://dc7b-160-72-87-42.ngrok.io/users", {
    method: 'POST',
    headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
    },
    body: `{
    "first_name": "${first}",
    "last_name": "${last}",
    "email": "${email}",
    "phone_number": "${phone}",
    "organization": "${company}",
    "job_title": "${position}"
    }`,
    });

    response.json().then(data => {
    console.log(JSON.stringify(data));
    });
  };

  return (
    <div>
      <img src={sayyes} width="240" height="200"></img>
      {!submitted ? (
        <form className="form-container" onSubmit={handleSubmit}>
          <label>
            First Name:
            <input type="text" value={first} onChange={handleFirstChange} />
          </label>
          <label>
            Last Name:
            <input type="text" value={last} onChange={handleLastChange} />
          </label>
          <label>
            Email:
            <input type="email" value={email} onChange={handleEmailChange} />
          </label>
          <label>
            Phone:
            <input type="text" value={phone} onChange={handlePhoneChange} />
          </label>
          <label>
            Hired By:
            <input type="text" value={company} onChange={handleCompanyChange} />
          </label>
          <label>
            Position:
            <input type="text" value={position} onChange={handlePositionChange} />
          </label>
          <button type="submit">Submit</button>
        </form>
      ) : (
        <div className="thank-you-message">Thank you for your submission!</div>
      )}
    </div>
  );
};

export default Form;