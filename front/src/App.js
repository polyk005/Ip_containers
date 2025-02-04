import React from 'react';
import ContainerTable from './components/ContainerTable';

const App = () => {
  return (
    <div style={{ padding: '20px' }}>
      <h1>Docker Container Monitoring</h1>
      <ContainerTable />
    </div>
  );
};

export default App;