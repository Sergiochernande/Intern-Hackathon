import './App.css';
import Form from './Form';

function App() {
<script>
      // Create a button element
      const button = document.createElement('button')

      // Set the button text to 'Can you click me?'
      button.innerText = 'I got the job'

      button.id = 'mainButton'

      // Attach the "click" event to your button
      button.addEventListener('click', () => {
        // When there is a "click"
        // it shows an alert in the browser
        alert('Oh, you clicked me!')
      })

      document.body.appendChild(button)
    </script>
  
  return (
    <div className="App">
      <header className="App-header">
        <Form/>
      </header>
    
        <body>
   <button id="mainButton">Click Me</button>
   <p id="count"></p>
   <script>
      var count = 0;
      var button = document.getElementById("mainButton");
      var countDisplay = document.getElementById("count");
      button.addEventListener("click", function() {
         count++;
         countDisplay.innerHTML = count;
      });
   </script>
</body>
    </div>
  );
}

export default App;
