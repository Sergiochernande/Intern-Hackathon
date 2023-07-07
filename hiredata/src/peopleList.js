import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './peopleList.css'
import sayyes from './images/sayyes.jpg'

const PeopleList = () => {
  const [applicants, setApplicants] = useState([]);
  const [sortedApplicants, setSortedApplicants] = useState([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [jobFilter, setJobFilter] = useState('');

  useEffect(() => {
    // Fetch people data from the API
    axios.get('https://dc7b-160-72-87-42.ngrok.io/users')
      .then(response => {
        setApplicants(response.data.applicants);
        setSortedApplicants(response.data.applicants);
      })
      .catch(error => {
        console.log(error);
      });
  }, []);

  console.log(sortedApplicants.length)

  const handleSort = () => {
    const sorted = [...sortedApplicants];
    sorted.sort((a, b) => (a.organization > b.organization) ? 1 : -1);
    setSortedApplicants(sorted);
  };

  const handleSearch = () => {
    if (Array.isArray(applicants)) {
      const filtered = applicants.filter(applicants =>
        applicants.first_name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        applicants.organization.toLowerCase().includes(searchQuery.toLowerCase()) ||
        applicants.job_title.toLowerCase().includes(searchQuery.toLowerCase()) ||
        applicants.last_name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        applicants.email.toLowerCase().includes(searchQuery.toLowerCase()) ||
        applicants.phone_number.toLowerCase().includes(searchQuery.toLowerCase())
      );
      setSortedApplicants(filtered);
    }
  };

  const handleReset = () => {
    setSearchQuery('');
    setSortedApplicants(applicants);
  };

  return (
    <div>
        <img src={sayyes} width="240" height="170"></img>
      <div>
        <input
          type="text"
          placeholder="Search by name, email, phone number, position or company"
          value={searchQuery}
          onChange={e => setSearchQuery(e.target.value)}
        />
        <button onClick={handleSearch}>Search</button>
        <button onClick={handleReset}>Reset</button>
      </div>

      <button onClick={handleSort}>Sort by Company</button>

      <ul>
        {sortedApplicants.length > 0 ? (
          sortedApplicants.map(applicants => (
            <li className='list-style'>
              <div>{applicants.first_name + " " + applicants.last_name}</div>
              <div>Email: {applicants.email}</div>
              <div>Phone Number: {applicants.phone_number}</div>
              <div>Company: {applicants.organization}</div>
              <div>Position: {applicants.job_title}</div>
            </li>
          ))
        ) : (
          <li>No people found.</li>
        )}
      </ul>
    </div>
  );
};

export default PeopleList;