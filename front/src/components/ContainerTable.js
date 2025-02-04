import React, { useEffect, useState } from 'react';
import { Table } from 'antd';
import { getContainers } from '../services/api';

const ContainerTable = () => {
  const [containers, setContainers] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const data = await getContainers();
      setContainers(data);
    };
  
    fetchData();
    const interval = setInterval(fetchData, 15000);
  
    return () => clearInterval(interval);
  }, []);

  const columns = [
    {
      title: 'IP Address',
      dataIndex: 'ip_address',
      key: 'ip_address',
    },
    {
      title: 'Ping Time',
      dataIndex: 'ping_time',
      key: 'ping_time',
    },
    {
      title: 'Last Successful Ping',
      dataIndex: 'last_successful_ping',
      key: 'last_successful_ping',
    },
  ];

  return (
    <Table
      dataSource={containers}
      columns={columns}
      rowKey="id"
      pagination={false}
    />
  );
};

export default ContainerTable;