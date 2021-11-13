package compiler

func importAll(code string) string {
	return importString + code + utilities
}

var importString = `
import com.sun.jarsigner.*;
import com.sun.java.accessibility.util.*;
import com.sun.jdi.*;
import com.sun.jdi.connect.*;
import com.sun.jdi.connect.spi.*;
import com.sun.jdi.event.*;
import com.sun.jdi.request.*;
import com.sun.management.*;
import com.sun.net.httpserver.*;
import com.sun.net.httpserver.spi.*;
import com.sun.nio.sctp.*;
import com.sun.security.auth.*;
import com.sun.security.auth.callback.*;
import com.sun.security.auth.login.*;
import com.sun.security.auth.module.*;
import com.sun.security.jgss.*;
import com.sun.source.doctree.*;
import com.sun.source.tree.*;
import com.sun.source.util.*;
import com.sun.tools.attach.*;
import com.sun.tools.attach.spi.*;
import com.sun.tools.javac.*;
import com.sun.tools.jconsole.*;
import java.applet.*;
import java.awt.*;
import java.awt.color.*;
import java.awt.datatransfer.*;
import java.awt.desktop.*;
import java.awt.dnd.*;
import java.awt.event.*;
import java.awt.font.*;
import java.awt.geom.*;
import java.awt.im.*;
import java.awt.im.spi.*;
import java.awt.image.*;
import java.awt.image.renderable.*;
import java.awt.print.*;
import java.beans.*;
import java.beans.beancontext.*;
import java.io.*;
import java.lang.*;
import java.lang.annotation.*;
import java.lang.constant.*;
import java.lang.instrument.*;
import java.lang.invoke.*;
import java.lang.management.*;
import java.lang.module.*;
import java.lang.ref.*;
import java.lang.reflect.*;
import java.math.*;
import java.net.*;
import java.net.http.*;
import java.net.spi.*;
import java.nio.*;
import java.nio.channels.*;
import java.nio.channels.spi.*;
import java.nio.charset.*;
import java.nio.charset.spi.*;
import java.nio.file.*;
import java.nio.file.attribute.*;
import java.nio.file.spi.*;
import java.rmi.*;
// import java.rmi.activation.*;
import java.rmi.dgc.*;
import java.rmi.registry.*;
import java.rmi.server.*;
import java.security.*;
import java.security.cert.*;
import java.security.interfaces.*;
import java.security.spec.*;
import java.sql.*;
import java.text.*;
import java.text.spi.*;
import java.time.*;
import java.time.chrono.*;
import java.time.format.*;
import java.time.temporal.*;
import java.time.zone.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;
import java.util.concurrent.locks.*;
import java.util.function.*;
import java.util.jar.*;
import java.util.logging.*;
import java.util.prefs.*;
import java.util.regex.*;
import java.util.spi.*;
import java.util.stream.*;
import java.util.zip.*;
import javax.accessibility.*;
import javax.annotation.processing.*;
import javax.crypto.*;
import javax.crypto.interfaces.*;
import javax.crypto.spec.*;
import javax.imageio.*;
import javax.imageio.event.*;
import javax.imageio.metadata.*;
import javax.imageio.plugins.bmp.*;
import javax.imageio.plugins.jpeg.*;
import javax.imageio.plugins.tiff.*;
import javax.imageio.spi.*;
import javax.imageio.stream.*;
import javax.lang.model.*;
import javax.lang.model.element.*;
import javax.lang.model.type.*;
import javax.lang.model.util.*;
import javax.management.*;
import javax.management.loading.*;
import javax.management.modelmbean.*;
import javax.management.monitor.*;
import javax.management.openmbean.*;
import javax.management.relation.*;
import javax.management.remote.*;
import javax.management.remote.rmi.*;
import javax.management.timer.*;
import javax.naming.*;
import javax.naming.directory.*;
import javax.naming.event.*;
import javax.naming.ldap.*;
import javax.naming.ldap.spi.*;
import javax.naming.spi.*;
import javax.net.*;
import javax.net.ssl.*;
import javax.print.*;
import javax.print.attribute.*;
import javax.print.attribute.standard.*;
import javax.print.event.*;
import javax.rmi.ssl.*;
import javax.script.*;
import javax.security.auth.*;
import javax.security.auth.callback.*;
import javax.security.auth.kerberos.*;
import javax.security.auth.login.*;
import javax.security.auth.spi.*;
import javax.security.auth.x500.*;
import javax.security.cert.*;
import javax.security.sasl.*;
import javax.sound.midi.*;
import javax.sound.midi.spi.*;
import javax.sound.sampled.*;
import javax.sound.sampled.spi.*;
import javax.sql.*;
import javax.sql.rowset.*;
import javax.sql.rowset.serial.*;
import javax.sql.rowset.spi.*;
import javax.swing.*;
import javax.swing.border.*;
import javax.swing.colorchooser.*;
import javax.swing.event.*;
import javax.swing.filechooser.*;
import javax.swing.plaf.*;
import javax.swing.plaf.basic.*;
import javax.swing.plaf.metal.*;
import javax.swing.plaf.multi.*;
import javax.swing.plaf.nimbus.*;
import javax.swing.plaf.synth.*;
import javax.swing.table.*;
import javax.swing.text.*;
import javax.swing.text.html.*;
import javax.swing.text.html.parser.*;
import javax.swing.text.rtf.*;
import javax.swing.tree.*;
import javax.swing.undo.*;
import javax.tools.*;
import javax.transaction.xa.*;
import javax.xml.*;
import javax.xml.catalog.*;
import javax.xml.crypto.*;
import javax.xml.crypto.dom.*;
import javax.xml.crypto.dsig.*;
import javax.xml.crypto.dsig.dom.*;
import javax.xml.crypto.dsig.keyinfo.*;
import javax.xml.crypto.dsig.spec.*;
import javax.xml.datatype.*;
import javax.xml.namespace.*;
import javax.xml.parsers.*;
import javax.xml.stream.*;
import javax.xml.stream.events.*;
import javax.xml.stream.util.*;
import javax.xml.transform.*;
import javax.xml.transform.dom.*;
import javax.xml.transform.sax.*;
import javax.xml.transform.stax.*;
import javax.xml.transform.stream.*;
import javax.xml.validation.*;
import javax.xml.xpath.*;
import jdk.dynalink.*;
import jdk.dynalink.beans.*;
import jdk.dynalink.linker.*;
import jdk.dynalink.linker.support.*;
import jdk.dynalink.support.*;
import jdk.javadoc.doclet.*;
import jdk.jfr.*;
import jdk.jfr.consumer.*;
import jdk.jshell.*;
import jdk.jshell.execution.*;
import jdk.jshell.spi.*;
import jdk.jshell.tool.*;
import jdk.management.jfr.*;
import jdk.net.*;
import jdk.nio.*;
import jdk.security.jarsigner.*;
import netscape.javascript.*;
import org.ietf.jgss.*;
import org.w3c.dom.*;
import org.w3c.dom.bootstrap.*;
import org.w3c.dom.css.*;
import org.w3c.dom.events.*;
import org.w3c.dom.html.*;
import org.w3c.dom.ls.*;
import org.w3c.dom.ranges.*;
import org.w3c.dom.stylesheets.*;
import org.w3c.dom.traversal.*;
import org.w3c.dom.views.*;
import org.w3c.dom.xpath.*;
import org.xml.sax.*;
import org.xml.sax.ext.*;
import org.xml.sax.helpers.*;


`

var utilities = `


class Uhrzeit implements Comparable{
    private int std;      // 0 bis 23
    private int min, sek; // 0 bis 59
    private int zone;     // Zeitzone in Werten von -11 bis +12
 
    private static int stdNormal;
    private static int minNormal;
 
    static void setzeNormalzeit (int std, int min)
    {
       stdNormal = std;
       minNormal = min;
    }
 
    Uhrzeit ()
    {
       std = stdNormal;
       min = minNormal;
       // weitere Angaben nicht n�tig (Standardvorbelegung)
    }
 
    Uhrzeit (int std, int min)
    {
       setzeStdMin (std, min);
    }
 
    Uhrzeit (int std, int min, int sek)
    {
       setzeUhrzeit (std, min, sek);
    }
 
    Uhrzeit (int std, int min, int sek, int zone)
    {
       setzeUhrzeit (std, min, sek, zone);
    }
 
    private void setzeStdMin (int std, int min)
    {
       if (min<0)
          throw new IllegalArgumentException
                    ("Minuten m�ssen >= 0 sein.");
       else
          while (min>=60)
             min = min - 60;
       this.min = min;
       if (std<0)
          throw new IllegalArgumentException
                    ("Stunden m�ssen >= 0 sein.");
       else
          while (std>=24)
             std = std - 24;
       this.std = std;
    }
 
    void setzeUhrzeit (int std, int min)
    {
       setzeStdMin (std, min);
       this.sek = 0;
    }
 
    void setzeUhrzeit (int std, int min, int sek)
    {
       setzeStdMin (std, min);
       if (sek<0 | sek>59)
          throw new IllegalArgumentException
             ("Sekunden m�ssen zwischen 0 und 59 liegen.");
       this.sek = sek;
       zone = 1;
    }
 
    void setzeUhrzeit (int std, int min, int sek, int zone)
    {
       setzeStdMin (std, min);
       if (sek<0 | sek>59)
          throw new IllegalArgumentException
             ("Sekunden m�ssen zwischen 0 und 59 liegen.");
       this.sek = sek;
       if (zone<-11 | zone>12)
          throw new IllegalArgumentException
             ("Zeitzone muss zwischen -11 und +12 liegen.");
       this.zone = zone;
    }
 
    String differenz(Uhrzeit bis)
    {
       int dStd = bis.std + bis.zone - std - zone;
       int dMin = bis.min - min;
       if (dMin < 0)
       {
          dMin = dMin + 60;
          dStd--;
       }
       int dSek = bis.sek - sek;
       if (dSek < 0)
       {
          dSek = dSek + 60;
          dMin--;
       }
       return dStd + " Std " + dMin + " Min " + dSek + " Sek";
    }
 
    public String toString ()
    {
       String zeit = ((std<10)?"0"+std:""+std) + ":";
       zeit = zeit + ((min<10)?"0"+min:""+min) + ":";
       zeit = zeit + ((sek<10)?"0"+sek:""+sek) + " GMT";
       zeit = zeit + ((zone<0)?"-":"+") + Math.abs(zone);
       return zeit;
    }
 
    public boolean equals(Object einObject)
    {
       Uhrzeit eineZeit = (Uhrzeit)einObject;
       return this.std - this.zone == eineZeit.std - eineZeit.zone
            & this.min == eineZeit.min;
    }
 
    public int compareTo(Object einObject)
    {
       Uhrzeit z = (Uhrzeit)einObject;
       return
          (((this.std-this.zone)*60)+this.min)*60+this.sek
         -((((z.std-z.zone)*60)+z.min)*60+z.sek);
    }
 
    int gibStunde ()
    {
       return std;
    }
 
    int gibMinute()
    {
       return min;
    }
 
 }

 class TastaturEingabe
 {
    static BufferedReader eing =
       new BufferedReader(new InputStreamReader (System.in));
 
    public static double readDouble (String prompt)
    {
       while (true)
       {
          System.out.print(prompt);
          try
          {
             String zeile = eing.readLine();
             return Double.parseDouble(zeile);
          }
          catch (NumberFormatException e)
          {
             System.out.println ("Bitte eine Zahl eingeben (ggf. mit Dezimalpunkt).");
          }
          catch (IOException e)
          {
             e.printStackTrace();
             System.exit(1);
          }
       }
    }
 
    public static int readInt (String prompt)
    {
       while (true)
       {
          System.out.print(prompt);
          try
          {
             String zeile = eing.readLine();
             return Integer.parseInt(zeile);
          }
          catch (NumberFormatException e)
          {
             System.out.println ("Bitte eine ganze Zahl eingeben.");
          }
          catch (IOException e)
          {
             e.printStackTrace();
             System.exit(1);
          }
       }
    }
 
    public static long readLong (String prompt)
    {
       while (true)
       {
          System.out.print(prompt);
          try
          {
             String zeile = eing.readLine();
             return Long.parseLong(zeile);
          }
          catch (NumberFormatException e)
          {
             System.out.println ("Bitte eine ganze Zahl eingeben.");
          }
          catch (IOException e)
          {
             e.printStackTrace();
             System.exit(1);
          }
       }
    }
 
    public static boolean readBoolean (String prompt)
    {
       while (true)
       {
          System.out.print(prompt);
          try
          {
             String e = eing.readLine();
             if (e.equals("j") || e.equals("J"))
                return true;
             if (e.equals("n") || e.equals("N"))
                return false;
             throw new IllegalArgumentException();
          }
          catch (IllegalArgumentException e)
          {
             System.out.println ("Bitte J oder N eingeben.");
          }
          catch (IOException e)
          {
             e.printStackTrace();
             System.exit(1);
          }
       }
    }
 
    public static String readString (String prompt)
    {
       while (true)
       {
          System.out.print(prompt);
          try
          {
             return eing.readLine();
          }
          catch (IOException e)
          {
             e.printStackTrace();
             System.exit(1);
          }
       }
    }
 
    public static char readChar (String prompt)
    {
       while (true)
       {
          System.out.print(prompt);
          try
          {
             String s = eing.readLine();
             if (s.length() == 0)
                throw new IllegalArgumentException();
             return s.charAt(0);
          }
          catch (IllegalArgumentException e)
          {
             System.out.println ("Bitte ein Zeichen eingeben.");
          }
          catch (IOException e)
          {
             e.printStackTrace();
             System.exit(1);
          }
       }
    }
 
 
    /**
     *  wartet auf <Return>
     *  @param b true: Meldung "Weiter ..." wird ausgegeben
     */
    public static void warte(boolean b)
    {
       String s;
       if (b)
       {
          System.out.print("....... Weiter mit <Return>");
          System.out.flush();
       }
       s = readString("");
    }
 
    /**
     *  wartet auf <Return>
     *  Meldung "....... Weiter mit <Return>" wird immer ausgegeben
     */
    public static void warte()
    {
       warte(true);
    }
 
    public static Datum readDatum (String prompt)
    {
       while (true)
       {
          System.out.print(prompt);
          try
          {
             String zeile = eing.readLine();
             if (zeile.length()!=10) throw new NumberFormatException();
             int tag = Integer.parseInt(zeile.substring(0,2));
             char trenn = zeile.charAt(2);
             if (trenn!='.' & trenn!='/') throw new NumberFormatException();
             int monat = Integer.parseInt(zeile.substring(3,5));
             trenn = zeile.charAt(5);
             if (trenn!='.' & trenn!='/') throw new NumberFormatException();
             int jahr = Integer.parseInt(zeile.substring(6,10));
             Datum einDatum = new Datum (tag, monat, jahr);
             return einDatum;
          }
          catch (NumberFormatException e)
          {
             System.out.println ("Bitte eine Datum in der Form tt.mm.jjjj eingeben.");
          }
          catch (IllegalArgumentException e)
          {
             System.out.println ("Fehler in Datum: " + e.getMessage());
          }
          catch (IOException e)
          {
             e.printStackTrace();
             System.exit(1);
          }
       }
    }
 
 }
 /**
 * Klasse zur formatierten Ausgabe einfacher Datentypen
 */

class FormatierteAusgabe
{
   private static int zehn_hoch(int potenz)
   {  int erg=1;
      for (int i=0;i<potenz;i++)
         erg=10*erg;
      return erg;
   }

/**
*  Formatierte Ausgabe ganzer Zahlen.
*  Es wird ausser der auszugebenden Zahl auch die Ausgabebreite
*  angegeben.
*  Die Zahl wird rechtsb�ndig in der angegebenen Breite ausgegeben.
*  Bei negativen Zahlen kommt das Minuszeichen unmittelbar vor
*  der ersten Ziffer.
*  Ist die Zahl zu gro�, wird sie unformatiert ausgegeben,
*  eingeleitet durch "|"
*  @param zahl auszugebende Zahl
*  @param breite Ausgabebreite (Anzahl auszugebender Zeichen)
*  @param fuell F�llzeichen
*/
   public static String format(int zahl,int breite,char fuell)
   {  boolean neg=false;
      String erg="";
      int lbreite=breite;
      if (zahl<0)
      {  neg=true;
         zahl=-zahl;
         lbreite--;
      }
      if (zahl>=zehn_hoch(lbreite))
      {  erg="|"+zahl;
         return erg;
      }
      for (int i=lbreite-1;i>0;i--)
         if (zahl<zehn_hoch(i))
            erg+=fuell; //System.out.print(" ");
         else
            break;
      if (neg)
         erg+="-";
      return erg+zahl;
   }

/**
*  Formatierte Ausgabe ganzer long-Zahlen.
*  Es wird ausser der auszugebenden Zahl auch die Ausgabebreite
*  angegeben.
*  Die Zahl wird rechtsb�ndig in der angegebenen Breite ausgegeben.
*  Bei negativen Zahlen kommt das Minuszeichen unmittelbar vor
*  der ersten Ziffer.
*  Ist die Zahl zu gro�, wird sie unformatiert ausgegeben,
*  eingeleitet durch "|"
*  @param zahl auszugebende Zahl
*  @param breite Ausgabebreite (Anzahl auszugebender Zeichen)
*  @param fuell F�llzeichen
*/
   public static String format(long zahl,int breite,char fuell)
   {  boolean neg=false;
      String erg="";
      int lbreite=breite;
      if (zahl<0)
      {  neg=true;
         zahl=-zahl;
         lbreite--;
      }
      if (zahl>=zehn_hoch(lbreite))
      {  erg="|"+zahl;
         return erg;
      }
      for (int i=lbreite-1;i>0;i--)
         if (zahl<zehn_hoch(i))
            erg+=fuell;
         else
            break;
      if (neg)
         erg+="-";
      return erg+zahl;
   }

/**
*  Formatierte Ausgabe ganzer int-Zahlen.
*  Es wird au�er der auszugebenden Zahl auch die Ausgabebreite
*  angegeben.
*  Die Zahl wird rechtsb�ndig in der angegebenen Breite ausgegeben.
*  Bei negativen Zahlen kommt das Minuszeichen unmittelbar vor
*  der ersten Ziffer.
*  Ist die Zahl zu gro�, wird sie unformatiert ausgegeben,
*  eingeleitet durch "|"
*  @param zahl auszugebende Zahl
*  @param breite Ausgabebreite (Anzahl auszugebender Zeichen)
*/
   public static void ausgabeInt(int zahl,int breite)
   {  System.out.print(format(zahl,breite,' ')); }

/**
*  Formatierte Ausgabe eines Strings.
*  Es wird ausser dem auszugebenden String auch die Ausgabebreite
*  angegeben.
*  Der String wird rechtsb�ndig ausgegeben und links mit einem
*  angebbaren Zeichen aufgef�llt.
*  @param s auszugebender String
*  @param breite Ausgabebreite (Anzahl auszugebender Zeichen)
*  @param fuell F�llzeichen
*/
   public static String format(String s,int breite,char fuell)
   {  int lg=s.length(),i;
      String erg=new String();
      if (s.length()>breite)
         return s.substring(s.length()-breite);
      for (i=0;i<breite-lg;i++)
         erg+=fuell;
      erg+=s;
      return erg;
   }



   final static long EINS = 1L;

   /**
   *  gibt eine long-Zahl als Bin�rzahl aus.
   *  nach jeweils 4 Bitstellen wird ein Zwischenraum ausgegeben;
   *  die Anzahl der ausgegebenen Bit ist angebbar
   *  @param zahl auszugebende long-Zahl
   *  @param breite Stellenanzahl f�r Bin�rzahl
   */
   public static void binaer(long zahl, int breite)
   {
      for (int i = breite-1; i>=0; i--)
      {  if ((zahl & (EINS<<i)) != 0)   //  Bitstelle berechnen
            System.out.print('1');
         else
            System.out.print('0');
         if (i%4 == 0)
            System.out.print(' ');      // Zwischenraum
      }
      System.out.flush();
   }

   public static void binaerAusgabe(long zahl, int breite)
   {
      binaer(zahl, breite);
   }

}

class Datum
{
   private int tag, monat, jahr;

   // aktuelles Datum (d.h., heute) erzeugen
   public Datum()
   {
      GregorianCalendar gc = new GregorianCalendar();
      tag = gc.get(GregorianCalendar.DAY_OF_MONTH);
      // In der Klasse GregorianCalendar werden die Monate von 0 bis 11
      // durchnummeriert; daher muessen wir den Wert 1 addieren.
      monat = gc.get(GregorianCalendar.MONTH) + 1;
      jahr = gc.get(GregorianCalendar.YEAR);
   }

   public Datum(int t, int m, int j)
   {
      tag = t;
      monat = m;
      jahr = j;
   }

   public String toString()
   {
      return tag + "." + monat + "." + jahr;
   }

   public boolean equals(Object obj)
   {
      if (obj instanceof Datum)
      {
         Datum d = (Datum)obj;

         return jahr == d.jahr
             && monat == d.monat
             && tag == d.tag;
      }
      else
         return false;
   }

   public boolean frueherAls(Datum d)
   {
      if(jahr < d.jahr)
         return true;
      if(jahr == d.jahr)
      {
         if(monat < d.monat)
            return true;
         if(monat == d.monat)
            return tag < d.tag;
      }
      return false;
   }

   public boolean vor(Datum d)
   {
      return frueherAls(d);
   }

   public boolean nach(Datum d)
   {
      return !vor(d) && !equals(d);
   }

}




`
