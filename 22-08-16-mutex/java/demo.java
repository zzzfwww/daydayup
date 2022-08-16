import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class demo {
    public static void main(String[] args) throws Exception{
        ReadWriteLock rw = new ReentrantReadWriteLock();
        rw.readLock().lock();
        System.out.println("read lock");
        Thread thread = new Thread(
                ()->{
                    rw.writeLock().lock();
                    System.out.println("write lock");
                    try {
                        Thread.sleep(1000L);
                    }catch (Exception e){
                        e.printStackTrace();
                    }
                    rw.writeLock().unlock();
                    System.out.println("write unlock");
                }
        );
        thread.start();
        Thread.sleep(500L);
        rw.readLock().unlock();
        System.out.println("read unlock");
        Thread.sleep(5000L);
    }
}
