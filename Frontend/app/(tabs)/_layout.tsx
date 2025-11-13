import React from 'react';
import FontAwesome from '@expo/vector-icons/FontAwesome';
import { Tabs } from 'expo-router';

import Colors from '@/constants/Colors';

// You can explore the built-in icon families and icons on the web at https://icons.expo.fyi/
function TabBarIcon(props: {
  name: React.ComponentProps<typeof FontAwesome>['name'];
  color: string;
}) {
  return <FontAwesome size={24} style={{ marginBottom: -2 }} {...props} />;
}

export default function TabLayout() {
  const theme = Colors['dark'];

  return (
    <Tabs
    screenOptions={{
      tabBarActiveTintColor: theme.tabIconSelected,
      tabBarInactiveTintColor: theme.tabIconDefault,
      tabBarStyle: {
          backgroundColor: theme.tabBarBackground,
          borderTopColor: theme.borderTop,
      },
      headerStyle: {
          backgroundColor: theme.background,
      },
      headerTitleStyle: {
          color: theme.text,
      },
      headerTintColor: theme.tint,
      headerShown: false,
    }}>
      <Tabs.Screen
        name="index"
        options={{
          title: 'Day',
          tabBarIcon: ({ color }) => <TabBarIcon name="sun-o" color={color} />,
        }}
      />
      <Tabs.Screen
        name="week"
        options={{
          title: 'Week',
          tabBarIcon: ({ color }) => <TabBarIcon name="calendar" color={color} />,
        }}
      />
      <Tabs.Screen
        name="month"
        options={{
          title: 'Month',
          tabBarIcon: ({ color }) => <TabBarIcon name="calendar-o" color={color} />,
        }}
      />
      <Tabs.Screen
        name="year"
        options={{
          title: 'Year',
          tabBarIcon: ({ color }) => <TabBarIcon name="line-chart" color={color} />,
        }}
      />
    </Tabs>
  );
}
