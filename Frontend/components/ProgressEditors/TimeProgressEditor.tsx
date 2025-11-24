import React, {useMemo, useState} from 'react';
import {ActivityIndicator, Pressable, StyleSheet, Text, View} from 'react-native';
import {TimerPickerModal} from 'react-native-timer-picker';
import {useColorScheme} from '@/components/useColorScheme';
import Colors from '@/constants/Colors';

interface TimeProgressEditorProps {
  // Value comes as HH:mm:ss.mmm
  value: string;
  onChange: (next: string) => void;
  disabled?: boolean;
  loading?: boolean;
}

// Format picked time to HH:mm:ss.000
function formatTime({ hours = 0, minutes = 0, seconds = 0 }: { hours?: number; minutes?: number; seconds?: number; }): string {
  const h = String(hours).padStart(2, '0');
  const m = String(minutes).padStart(2, '0');
  const s = String(seconds).padStart(2, '0');
  return `${h}:${m}:${s}.000`;
}

export default function TimeProgressEditor({ value, onChange, disabled, loading }: TimeProgressEditorProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'dark'];

  const [showPicker, setShowPicker] = useState(false);
  const editable = useMemo(() => !disabled && !loading, [disabled, loading]);

  const summary = useMemo(() => (value ? value.substring(0, 8) : '00:00:00'), [value]);

  return (
    <View style={styles.container}>
      <Pressable
        testID="open-time-picker"
        accessibilityRole="button"
        accessibilityLabel="Edit time"
        accessibilityState={{ disabled: !editable, busy: !!loading }}
        disabled={!editable}
        onPress={() => setShowPicker(true)}
        style={({ pressed }) => [
          styles.button,
          { borderColor: theme.borderTop, backgroundColor: theme.componentBackground },
          (!editable) && { opacity: 0.7 },
          pressed && editable && { transform: [{ scale: 0.98 }] },
        ]}
      >
        <Text style={[styles.buttonText, { color: theme.text }]}>{summary}</Text>
        {loading && (
          <View style={styles.loaderOverlay} pointerEvents="none">
            <ActivityIndicator color={theme.text} size="small" />
          </View>
        )}
      </Pressable>

      <TimerPickerModal
        visible={showPicker}
        setIsVisible={setShowPicker}
        onConfirm={(picked) => {
          const formatted = formatTime(picked ?? {});
          onChange(formatted);
          setShowPicker(false);
        }}
        onCancel={() => setShowPicker(false)}
        closeOnOverlayPress
        styles={{ theme: 'light' }}
        // For time-of-day feel; leaving 24h by default, no AM/PM
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    width: '100%',
    alignItems: 'center',
    paddingTop: 8,
  },
  button: {
    minWidth: 140,
    height: 44,
    borderRadius: 8,
    borderWidth: 2,
    alignItems: 'center',
    justifyContent: 'center',
    paddingHorizontal: 12,
  },
  buttonText: {
    fontSize: 18,
    fontWeight: 'bold',
    textAlign: 'center',
  },
  loaderOverlay: {
    position: 'absolute',
    right: 12,
    height: 44,
    justifyContent: 'center',
  },
});
