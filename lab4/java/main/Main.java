package lab4.java.main;

import java.io.File;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;

public class Main {
    public static void main(String args[]) {
        System.out.println("hejka");
        String inputFileName = "test/example0.tga";
        String outputFileName = "output/example0.tga";
        if (args.length > 0) {
            inputFileName = args[0];
        }
        if (args.length > 1) {
            outputFileName = args[1];
        }
        File imageInput = new File(inputFileName);
        byte[] fileContent = null;
        try {
            fileContent = Files.readAllBytes(imageInput.toPath());
        } catch (IOException e) {
            System.out.println(e);
            return;
        }

        System.out.println(fileContent.length);
        for (int i = 0; i < 128; i++) {
            System.out.format("%d: %d\n", i, Integer.to(fileContent[i]));

            // System.out.format("%d: %8s\n", i, Integer.toBinaryString(fileContent[i] & 0xFF).replace(' ', '0'));
        }
    }
}