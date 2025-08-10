import React from 'react';
import { StyleSheet } from 'react-native';
import { Text, View } from '@/components/Themed';

export default function YearScreen() {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Year view</Text>
      <Text style={styles.caption}>Coming soon</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    gap: 8,
  },
  title: {
    fontSize: 20,
    fontWeight: 'bold',
  },
  caption: {
    opacity: 0.7,
  },
});
