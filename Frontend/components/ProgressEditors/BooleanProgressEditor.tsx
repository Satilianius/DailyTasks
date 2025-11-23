import {ActivityIndicator, Pressable, StyleSheet, Text, View} from "react-native";
import {useColorScheme} from '@/components/useColorScheme';
import Colors from '@/constants/Colors';

interface BooleanProgressEditorProps {
  value: boolean;
  onChange: (next: boolean) => void;
  disabled?: boolean;
  loading?: boolean;
}

export default function BooleanProgressEditor({ value, onChange, disabled, loading }: BooleanProgressEditorProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'dark'];

  return (
    <View style={styles.container}>
      <Pressable
        accessibilityRole="checkbox"
        accessibilityState={{ checked: value, disabled: !!disabled, busy: !!loading }}
        onPress={() => { if (!disabled && !loading) onChange(!value); }}
        style={({ pressed }) => [
          styles.checkbox,
          { borderColor: theme.borderTop, backgroundColor: theme.componentBackground },
          value && { backgroundColor: theme.successTint },
          (disabled || loading) && { opacity: 0.6 },
          pressed && !disabled && !loading && { transform: [{ scale: 0.90 }] },
        ]}
      >
        {loading ? (
          <ActivityIndicator color={theme.text} size="small" />
        ) : (
          <Text style={[styles.checkboxLabel, { color: value ? theme.background : theme.text }]}>
            {value ? '✓' : ''}
          </Text>
        )}
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    width: '100%',
    alignItems: 'center',
    paddingTop: 8,
  },
  checkbox: {
    width: 44,
    height: 44,
    borderRadius: 8,
    borderWidth: 2,
    alignItems: 'center',
    justifyContent: 'center',
  },
  checkboxLabel: {
    fontSize: 24,
    fontWeight: 'bold',
  },
});