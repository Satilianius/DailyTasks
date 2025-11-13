import {Pressable, StyleSheet} from "react-native";
import {Text, View} from "@/components/Themed";
import FontAwesome from '@expo/vector-icons/FontAwesome';

interface DateNavigatorProps {
  currentDate: Date;
  onPrevious: () => void;
  onNext: () => void;
}

export default function DateNavigator({currentDate, onPrevious, onNext}: DateNavigatorProps) {
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

      <Pressable onPress={onPrevious} style={styles.button}>
        <FontAwesome name="chevron-left" size={20} color="white"/>
      </Pressable>

      <View style={styles.dateDisplay}>
        <Text style={styles.buttonText}>{getCurrentDateText(currentDate)}</Text>
      </View>

      <Pressable onPress={onNext} style={styles.button}>
        <FontAwesome name="chevron-right" size={20} color="white" />
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
    backgroundColor: '#6d6d6d',
    alignItems: 'center',
    justifyContent: 'center',
  },
  buttonText: {
    color: '#FFF',
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
    backgroundColor: '#6d6d6d',
    marginHorizontal: 8,
    alignItems: 'center',
    justifyContent: 'center',
  },
});