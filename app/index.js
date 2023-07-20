import {View, Text, StyleSheet, TextInput, Button} from "react-native";
import {Link} from "expo-router";

export default function App() {
    return (
        <View style={styles.container}>
            <Text>Hennes Namn</Text>
            <TextInput autoComplete="name-given" style={styles.input}/>
            <Text>Hans namn</Text>
            <TextInput autoComplete="name-given" style={styles.input}/>
            <Link style={styles.input} href="/start">Starta</Link>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
        flexDirection: "column",
    },
    input: {
        width: 300,
        height: 40,
        margin: 12,
        borderWidth: 1,
        padding: 10,
    },
});