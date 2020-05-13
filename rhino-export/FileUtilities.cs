using System.IO;

namespace ExportAR
{
    public class FileUtilties

    {

        /// <summary>
        /// Deletes all files in the directory.
        /// </summary>
        /// <param name="dir">The directory to delete files in.</param>
        public static void DeleteFiles(string dir)
        {
            foreach (string file in Directory.EnumerateFiles(dir)) File.Delete(file);
        }

        public static string EncodeBase64(string path)
        {
            string plainText = File.ReadAllText(path);
            var plainTextBytes = System.Text.Encoding.UTF8.GetBytes(plainText);
            return System.Convert.ToBase64String(plainTextBytes);
        }

        public static void WriteToTxt(string fileName, string txt)
        {
            string path = fileName;
            if (!File.Exists(path))
            {
                using (StreamWriter sw = File.CreateText(path))
                {
                    sw.WriteLine(txt);
                }
            }
            else
            {
                using (StreamWriter sw = File.AppendText(path))
                {
                    sw.WriteLine(txt);
                }
            }
        }

    }
}