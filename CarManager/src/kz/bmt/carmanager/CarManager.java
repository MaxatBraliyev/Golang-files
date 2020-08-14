package kz.bmt.carmanager;

public class CarManager {

	public static void main(String[] args) throws Exception {

		Car prius = new Car("Toyota Prius", CarColor.BLACK, 2008, 15000, 1600);
		
		Car renault = new Car("Renaul Logan", CarColor.GREEN, 2012, 10000, 1800);
		
		Car renault2 = new Car("Renaul Logan", CarColor.GREEN, 2012, 10000, 1800);
		
		prius.addCarDistance(5);
		prius.addCarDistance(20.5);
		
		System.out.println(prius);
		System.out.println(renault);
		System.out.println(renault2);
		System.out.println(prius.equals(renault));
		System.out.println(renault.equals(renault2));
	}

}
