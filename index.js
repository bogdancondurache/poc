import React from 'react';
import ReactDOM from 'react-dom';
import { ThemeProvider, Button } from '@robocorp/ds';

function App() {
  const value = document.variable
  return (
    <ThemeProvider theme="dark">
      <Button variant="primary">{value}</Button>
    </ThemeProvider>
  );
}

ReactDOM.render(<App />, document.querySelector('#app'));