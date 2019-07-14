/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import {storiesOf} from '@storybook/react';
import AppSideBar from '../components/layout/AppSideBar.react';
import AssignmentIcon from '@material-ui/icons/Assignment';
import ListIcon from '@material-ui/icons/List';
import React, {useState} from 'react';
import SettingsIcon from '@material-ui/icons/Settings';

const ExpandablePanel = () => {
  const [isExpanded, setIsExpanded] = useState();
  const [showExpandButton, setShowExpandButton] = useState(false);
  return (
    <div style={{display: 'flex'}}>
      <AppSideBar
        user={{email: 'user@fb.com', isSuperUser: false}}
        projects={[]}
        mainItems={[
          <AssignmentIcon color="primary" style={{margin: '8px'}} />,
          <ListIcon color="primary" style={{margin: '8px'}} />,
        ]}
        secondaryItems={[]}
        useExpandButton={true}
        showExpandButton={showExpandButton}
        onExpandClicked={() => setIsExpanded(!isExpanded)}
        expanded={isExpanded}
      />
      <div
        onMouseEnter={() => setShowExpandButton(true)}
        onMouseLeave={() => setShowExpandButton(false)}
        style={{
          height: '100vh',
          width: '300px',
          backgroundColor: 'blue',
          visibility: isExpanded ? 'visible' : 'hidden',
        }}
      />
    </div>
  );
};

storiesOf('AppSideBar', module)
  .add('default', () => (
    <AppSideBar
      user={{email: 'user@fb.com', isSuperUser: true}}
      projects={[
        {
          id: 'inventory',
          name: 'Inventory',
          secondary: 'Inventory Management',
          url: '/',
        },
        {
          id: 'nms',
          name: 'NMS',
          secondary: 'Network Management',
          url: '/nms',
        },
      ]}
      mainItems={[
        <AssignmentIcon color="primary" style={{margin: '8px'}} />,
        <ListIcon color="primary" style={{margin: '8px'}} />,
      ]}
      secondaryItems={[
        <SettingsIcon color="primary" style={{margin: '8px'}} />,
      ]}
    />
  ))
  .add('Not super user', () => (
    <AppSideBar
      user={{email: 'user@fb.com', isSuperUser: false}}
      projects={[
        {
          id: 'inventory',
          name: 'Inventory',
          secondary: 'Inventory Management',
          url: '/',
        },
        {
          id: 'nms',
          name: 'NMS',
          secondary: 'Network Management',
          url: '/nms',
        },
      ]}
      mainItems={[
        <AssignmentIcon color="primary" style={{margin: '8px'}} />,
        <ListIcon color="primary" style={{margin: '8px'}} />,
      ]}
      secondaryItems={[]}
    />
  ))
  .add('No projects', () => (
    <AppSideBar
      user={{email: 'user@fb.com', isSuperUser: false}}
      projects={[]}
      mainItems={[
        <AssignmentIcon color="primary" style={{margin: '8px'}} />,
        <ListIcon color="primary" style={{margin: '8px'}} />,
      ]}
      secondaryItems={[]}
    />
  ))
  .add('Expand Button', () => <ExpandablePanel />);
