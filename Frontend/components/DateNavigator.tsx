import {Pressable, StyleSheet} from "react-native";
import {View, Text} from "@/components/Themed";

interface DateNavigatorProps {
    currentDate: Date;
    onPrevious: () => void;
    onNext: () => void;
}

export default function DateNavigator({currentDate, onPrevious, onNext} : DateNavigatorProps) {
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
                <Text style={styles.buttonText}>{'<'}</Text>
            </Pressable>

            <View style={styles.dateDisplay}>
                <Text>{getCurrentDateText(currentDate)}</Text>
            </View>

            <Pressable onPress={onNext} style={styles.button}>
                <Text style={styles.buttonText}>{'>'}</Text>
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
        fontWeight: 'bold',
        padding: 8,
    },
    dateDisplay: {
        borderRadius: 8,
        height: 40,
        width: 100,
        paddingHorizontal:8,
        flex: 1,
        flexDirection: 'row',
        backgroundColor: '#6d6d6d',
        marginHorizontal: 8,
        alignItems: 'center',
        justifyContent: 'center',
    },
});