import {Pressable, StyleSheet} from "react-native";
import {Text, View} from "@/components/Themed";
import FontAwesome from '@expo/vector-icons/FontAwesome';
import Colors from "@/constants/Colors";
import { useColorScheme } from "@/components/useColorScheme";

interface DateNavigatorProps {
  currentDate: Date;
  onPrevious: () => void;
  onNext: () => void;
}

export default function DateNavigator({currentDate, onPrevious, onNext}: DateNavigatorProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'dark'];

  function getCurrentDateText(currentDate: Date) {
    let today = new Date();
    let yesterday = new Date(today.getTime())
    yesterday.setDate(today.getDate() - 1);

    return currentDate.toDateString() === today.toDateString()
      ? 'Today'
      : currentDate.toDateString() === yesterday.toDateString()
        ? 'Yesterday'
        : currentDate.toLocaleDateString();
  }

  return (
    <View style={styles.container}>

      <Pressable onPress={onPrevious} style={[styles.button, { backgroundColor: theme.componentBackground }]}>
        <FontAwesome name="chevron-left" size={20} color={theme.text}/>
      </Pressable>

      <View style={[styles.dateDisplay, { backgroundColor: theme.componentBackground }]}>
        <Text style={[styles.buttonText, { color: theme.text }]}>{getCurrentDateText(currentDate)}</Text>
      </View>

      <Pressable onPress={onNext} style={[styles.button, { backgroundColor: theme.componentBackground }]}>
        <FontAwesome name="chevron-right" size={20} color={theme.text} />
      </Pressable>
    </View>);
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingHorizontal: 16,
    paddingVertical: 12,
  },
  button: {
    borderRadius: 8,
    width: 50,
    height: 40,
    alignItems: 'center',
    justifyContent: 'center',
  },
  buttonText: {
    fontSize: 20,
    padding: 8,
  },
  dateDisplay: {
    borderRadius: 8,
    height: 40,
    minWidth: 150,
    paddingHorizontal: 8,
    flex: 1,
    flexDirection: 'row',
    marginHorizontal: 8,
    alignItems: 'center',
    justifyContent: 'center',
  },
});