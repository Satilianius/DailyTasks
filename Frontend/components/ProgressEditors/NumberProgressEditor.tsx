import React, {useEffect, useMemo, useState} from 'react';
import {ActivityIndicator, StyleSheet, TextInput, View} from 'react-native';
import {useColorScheme} from '@/components/useColorScheme';
import Colors from '@/constants/Colors';

interface NumberProgressEditorProps {
  value: number;
  onChange: (next: number) => void;
  disabled?: boolean;
  loading?: boolean;
}

export default function NumberProgressEditor({ value, onChange, disabled, loading }: NumberProgressEditorProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'dark'];

  const [text, setText] = useState<string>(() => String(value ?? 0));

  useEffect(() => {
    // Keep the local state in sync when the external value changes (e.g. optimistic updates)
    setText(String(value ?? 0));
  }, [value]);

  const editable = useMemo(() => !disabled && !loading, [disabled, loading]);

  const handleChangeText = (next: string) => {
    // Normalize comma to dot for decimals
    const normalized = next.replace(',', '.');

    // Allow typing of a few transient-but-valid prefixes while editing
    // (empty, just sign, just dot, or sign + dot). These don't emit onChange.
    const transient = ['', '+', '-', '.', '+.', '-.'];
    if (transient.includes(normalized)) {
      setText(next);
      return;
    }

    // Fully editable pattern: only digits with an optional single leading sign and single dot
    const editablePattern = /^[-+]?\d*\.?\d*$/;
    if (!editablePattern.test(normalized)) {
      // Reject any input containing letters or other disallowed characters by ignoring the change
      return;
    }

    // Accept the visual change
    setText(next);

    // If it parses to a number, emit it upstream
    const num = Number(normalized);
    if (!Number.isNaN(num)) {
      onChange(num);
    }
  };

  const handleEndEditing = () => {
    // If input is empty or invalid, revert to the last valid value
    if (text.trim() === '' || Number.isNaN(Number(text.replace(',', '.')))) {
      setText(String(value ?? 0));
    }
  };

  return (
    <View style={styles.container}>
      <TextInput
        value={text}
        onChangeText={handleChangeText}
        onEndEditing={handleEndEditing}
        editable={editable}
        keyboardType="numeric"
        style={[
          styles.input,
          { borderColor: theme.borderTop, color: theme.text, backgroundColor: theme.componentBackground },
          (!editable) && { opacity: 0.7 },
        ]}
        placeholder="Enter number"
        placeholderTextColor={theme.text}
      />
      {loading && (
        <View style={styles.loaderOverlay} pointerEvents="none">
          <ActivityIndicator color={theme.text} size="small" />
        </View>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    width: '100%',
    alignItems: 'center',
    paddingTop: 8,
  },
  input: {
    width: 120,
    height: 44,
    borderRadius: 8,
    borderWidth: 2,
    fontSize: 18,
    fontWeight: 'bold',
    textAlign: 'center',
    paddingHorizontal: 8,
  },
  loaderOverlay: {
    position: 'absolute',
    right: 12,
    height: 44,
    justifyContent: 'center',
  },
});
