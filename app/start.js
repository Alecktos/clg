import {Button, StyleSheet, Text, TextInput, View} from "react-native";
import {Link} from "expo-router";

export default function Start() {
    return (
        <View style={styles.container}>
            <Text>Nu är jag på Start Sidan</Text>
            <Link style={styles.input} href="/">Tillbaka</Link>
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